package main

type Asteroid struct {
	Id       string   `json:"i"`
	Alive    bool     `json:"z"`
	Died     uint64   `json:"-"`
	Position *Point   `json:"p"`
	Angle    float64  `json:"a"`
	Velocity *Vector  `json:"v"`
	Shape    []*Point `json:"s"`
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
		RoundVector(AngleAndSpeedToVector(RandomAngle(), float64(Random(10, 50)))),
		shape
}

func CreateAsteroid(id string, p *Point, a float64, v *Vector, s []*Point) *Asteroid {
	return &Asteroid{
		Id:       id,
		Alive:    true,
		Died:     0,
		Position: p,
		Angle:    a,
		Velocity: v,
		Shape:    s,
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
	}
}
