package main

type Asteroid struct {
	Id       string        `json:"id"`
	Position *Coordinate   `json:"position"`
	Angle    Angle         `json:"angle"`
	Velocity uint16        `json:"velocity"`
	Shape    []*Coordinate `json:"shape"`
}

func (a *Asteroid) Tick(t uint64) {
	x, y := AmountToMove(a.Angle, a.Velocity, t)
	a.Position.X += x
	a.Position.Y += y
}

func CreateAsteroid(id string) *Asteroid {
	sides := Random(6, 9)
	shape := make([]*Coordinate, sides)
	shape[0] = &Coordinate{0, 0}
	last := shape[0]
	totalAngle := 0
	for i := 1; i < sides; i++ {
		a := (360 - totalAngle) / (sides - i + 1)
		totalAngle += Random(a-5, a+5)
		l := Random(8, 15)
		v := AngleToVector(Angle(totalAngle))
		c := &Coordinate{last.X + int64(v.X*float64(l)), last.Y + int64(v.Y*float64(l))}
		shape[i] = c
		last = c
	}

	return &Asteroid{
		Id:       id,
		Angle:    RandomAngle(),
		Position: &Coordinate{int64(Random(-1000, 1000)), int64(Random(-1000, 1000))},
		Velocity: uint16(Random(10, 50)),
		Shape:    shape,
	}
}
