package main

type Asteroid struct {
	Id       string   `json:"i"`
	Alive    bool     `json:"z"`
	Died     uint64   `json:"-"`
	Position *Point   `json:"p"`
	Angle    float64  `json:"a"`
	Velocity *Vector  `json:"v"`
	Shape    []*Point `json:"s"`
	Radius   float64  `json:"d"`
}

const (
	Small = iota
	Medium
	Large
)

func RandomAsteroidShape(size int) []*Point {
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
	shape := make([]*Point, sides)
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

func RandomAsteroidGeometry() (*Point, float64, *Vector, []*Point) {
	return MakePoint(float64(Random(-1000, 1000)), float64(Random(-1000, 1000))),
		RandomAngle(),
		RoundVector(AngleAndSpeedToVector(RandomAngle(), float64(Random(10, 50)))),
		RandomAsteroidShape(Random(Small, Large))
}

func CreateAsteroid(id string, position *Point, angle float64, velocity *Vector, shape []*Point) *Asteroid {
	// calculate radius
	var radius float64 = 0
	zero := &Point{0, 0}
	for _, p := range shape {
		d := DistanceBetweenPoints(zero, p)
		if d > radius {
			radius = d
		}
	}

	return &Asteroid{
		Id:       id,
		Alive:    true,
		Died:     0,
		Position: position,
		Angle:    angle,
		Velocity: velocity,
		Shape:    shape,
		Radius:   radius,
	}
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
		Id:       a.Id,
		Alive:    a.Alive,
		Died:     a.Died,
		Position: pos,
		Angle:    a.Angle,
		Velocity: a.Velocity,
		Shape:    a.Shape,
		Radius:   a.Radius,
	}
}
