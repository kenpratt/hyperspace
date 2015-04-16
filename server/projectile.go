package main

type Projectile struct {
	Id       string    `json:"id"`
	Position *Position `json:"position"`
	Angle    Angle     `json:"angle"`
}

func (p *Projectile) Tick(t uint64) {
	x, y := AmountToMove(p.Angle, game.constants.ProjectileSpeed, t)
	p.Position.X += x
	p.Position.Y += y
}
