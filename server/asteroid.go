package main

type Asteroid struct {
	Id       string    `json:"id"`
	Position *Position `json:"position"`
	Angle    Angle     `json:"angle"`
	Velocity uint16    `json:"velocity"`
	Width    uint8     `json:"width"`
}

func (a *Asteroid) Tick(t uint64) {
	x, y := AmountToMove(a.Angle, a.Velocity, t)
	a.Position.X += x
	a.Position.Y += y
}

func CreateAsteroid(id string) *Asteroid {
	return &Asteroid{
		Id:       id,
		Angle:    RandomAngle(),
		Position: &Position{int64(Random(-1000, 1000)), int64(Random(-1000, 1000))},
		Velocity: uint16(Random(20, 100)),
		Width:    uint8(Random(5, 10)),
	}
}
