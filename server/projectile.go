package main

type Projectile struct {
	Id       string  `json:"id"`
	Position *Point  `json:"position"`
	Velocity *Vector `json:"velocity"`
	Created  uint64  `json:"created"`
	Owner    string  `json:"owner"`
}

const (
	ProjectileRadius = 10
)

func (p *Projectile) Tick(t uint64) {
	x, y := AmountToMove(p.Velocity, t)
	p.Position.X += x
	p.Position.Y += y
}

func (p *Projectile) Alive() bool {
	return (MakeTimestamp() - p.Created) < 2000
}
