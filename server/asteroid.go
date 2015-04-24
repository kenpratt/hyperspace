package main

type Asteroid struct {
	Id       string   `bson:"i"`
	Alive    bool     `bson:"z"`
	Position *Point   `bson:"p"`
	Angle    float64  `bson:"a"`
	Velocity *Vector  `bson:"v"`
	Shape    []*Point `bson:"s"`
}

func RandomAsteroidGeometry() (*Point, float64, *Vector, []*Point) {
	sides := Random(6, 9)
	shape := make([]*Point, sides)
	shape[0] = MakePoint(0, 0)
	last := shape[0]
	totalAngle := 0.0
	for i := 1; i < sides; i++ {
		a := (360 - int(totalAngle)) / (sides - i + 1)
		totalAngle += float64(Random(a-5, a+5))
		l := Random(8, 15)
		v := AngleToVector(totalAngle)
		c := MakePoint(last.X+v.X*float64(l), last.Y+v.Y*float64(l))
		shape[i] = c
		last = c
	}

	return MakePoint(float64(Random(-1000, 1000)), float64(Random(-1000, 1000))),
		RandomAngle(),
		RoundVector(AngleAndSpeedToVector(RandomAngle(), uint16(Random(10, 50)))),
		shape
}

func CreateAsteroid(id string, p *Point, a float64, v *Vector, s []*Point) *Asteroid {
	return &Asteroid{
		Id:       id,
		Alive:    true,
		Position: p,
		Angle:    a,
		Velocity: v,
		Shape:    s,
	}
}

func (a *Asteroid) Tick(t uint64, state *GameState) *Asteroid {
	// calculate time since last update (in milliseconds)
	elapsed := t - state.Time

	// calculate new position
	x, y := AmountToMove(a.Velocity, elapsed)
	pos := MakePoint(a.Position.X+x, a.Position.Y+y)

	// return copy of object with new position
	return &Asteroid{
		Id:       a.Id,
		Alive:    a.Alive,
		Position: pos,
		Angle:    a.Angle,
		Velocity: a.Velocity,
		Shape:    a.Shape,
	}
}
