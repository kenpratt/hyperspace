package main

type Shape []*Point

type Asteroid struct {
	Id        string  `json:"i"`
	Alive     bool    `json:"z"`
	Died      uint64  `json:"-"`
	Size      int     `json:"-"`
	Position  *Point  `json:"p"`
	Angle     float64 `json:"a"`
	Velocity  *Vector `json:"v"`
	Shape     Shape   `json:"s"`
	SubShapes []Shape `json:"-"`
	Radius    float64 `json:"d"`
}

type AsteroidGeometry struct {
	Size      int
	Position  *Point
	Angle     float64
	Velocity  *Vector
	Shape     Shape
	SubShapes []Shape
}

const (
	Small = iota
	Medium
	Large
)

func RandomAsteroidShape(size int) Shape {
	var minSides, maxSides, targetCirumference int
	switch size {
	case Small:
		minSides = 6
		maxSides = 8
		targetCirumference = 45
	case Medium:
		minSides = 6
		maxSides = 9
		targetCirumference = 90
	case Large:
		minSides = 7
		maxSides = 10
		targetCirumference = 160
	}

	sides := Random(minSides, maxSides)

	minSideLength := float64(targetCirumference) / float64(sides) * 0.7
	maxSideLength := float64(targetCirumference) / float64(sides) * 1.3

	// randomly generate a shape
	shape := make(Shape, sides)
	shape[0] = MakePoint(0, 0)
	last := shape[0]
	totalAngle := 0.0
	for i := 1; i < sides; i++ {
		a := (360 - int(totalAngle)) / (sides - i + 1)
		totalAngle += float64(Random(a-5, a+5))
		l := RandomFloat(minSideLength, maxSideLength)
		v := AngleToVector(totalAngle)
		c := MakePoint(last.X+v.X*float64(l), last.Y+v.Y*float64(l))
		shape[i] = c
		last = c
	}

	// find center by averaging points, and normalize points around center
	center := CalculateCenter(shape)
	for _, c := range shape {
		c.X -= center.X
		c.Y -= center.Y
	}

	return shape
}

func RandomAsteroidGeometry() *AsteroidGeometry {
	size := Random(Small, Large)
	subShapeCount := 0
	switch size {
	case Large:
		subShapeCount = 6
	case Medium:
		subShapeCount = 2
	}

	subShapes := make([]Shape, subShapeCount, subShapeCount)
	for i, _ := range subShapes {
		if i < 2 {
			subShapes[i] = RandomAsteroidShape(size - 1)
		} else {
			subShapes[i] = RandomAsteroidShape(size - 2)
		}
	}

	return &AsteroidGeometry{
		Size:      size,
		Position:  MakePoint(float64(Random(-1000, 1000)), float64(Random(-1000, 1000))),
		Angle:     RandomAngle(),
		Velocity:  RoundVector(AngleAndSpeedToVector(RandomAngle(), float64(Random(10, 50)))),
		Shape:     RandomAsteroidShape(size),
		SubShapes: subShapes,
	}
}

func CreateAsteroid(id string, geom *AsteroidGeometry) *Asteroid {
	return &Asteroid{
		Id:        id,
		Alive:     true,
		Died:      0,
		Size:      geom.Size,
		Position:  geom.Position,
		Angle:     geom.Angle,
		Velocity:  geom.Velocity,
		Shape:     geom.Shape,
		SubShapes: geom.SubShapes,
		Radius:    geom.Shape.Radius(),
	}
}

func (s Shape) Radius() float64 {
	// calculate radius
	var radius float64 = 0
	zero := &Point{0, 0}
	for _, p := range s {
		d := DistanceBetweenPoints(zero, p)
		if d > radius {
			radius = d
		}
	}
	return radius
}

func (a *Asteroid) Tick(t uint64, state *GameState) *Asteroid {
	// calculate time since last update (in milliseconds)
	elapsedMillis := t - state.Time

	// elapsed time in percentage of a second
	elapsed := float64(elapsedMillis) / 1000

	// calculate new position
	pos := MakePoint(a.Position.X+a.Velocity.X*elapsed, a.Position.Y+a.Velocity.Y*elapsed)

	// return copy of object with new position
	return &Asteroid{
		Id:        a.Id,
		Alive:     a.Alive,
		Died:      a.Died,
		Size:      a.Size,
		Position:  pos,
		Angle:     a.Angle,
		Velocity:  a.Velocity,
		Shape:     a.Shape,
		SubShapes: a.SubShapes,
		Radius:    a.Radius,
	}
}

func (a *Asteroid) Splittable() bool {
	return a.Size > Small
}

func (a *Asteroid) Split() (*Asteroid, *Asteroid) {
	vector := AngleToVector(a.Angle)

	shape1 := a.SubShapes[0]
	shape2 := a.SubShapes[1]

	position1 := AddVectorToPoint(MultiplyVector(vector, shape1.Radius()), a.Position)
	position2 := AddVectorToPoint(MultiplyVector(vector, -shape1.Radius()), a.Position)

	velocity1 := AddVectors(a.Velocity, MultiplyVector(vector, 10))
	velocity2 := AddVectors(a.Velocity, MultiplyVector(vector, -10))

	var subShapes1, subShapes2 []Shape
	if a.Size == Large {
		subShapes1 = a.SubShapes[2:4]
		subShapes2 = a.SubShapes[4:6]
	}

	a1 := CreateAsteroid(
		a.Id+".1",
		&AsteroidGeometry{
			Size:      a.Size - 1,
			Position:  position1,
			Angle:     a.Angle,
			Velocity:  velocity1,
			Shape:     shape1,
			SubShapes: subShapes1,
		},
	)
	a2 := CreateAsteroid(
		a.Id+".2",
		&AsteroidGeometry{
			Size:      a.Size - 1,
			Position:  position2,
			Angle:     a.Angle,
			Velocity:  velocity2,
			Shape:     shape2,
			SubShapes: subShapes2,
		},
	)
	return a1, a2
}
