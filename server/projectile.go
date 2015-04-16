package main

type Projectile struct {
	Id       string    `json:"id"`
	Position *Position `json:"position"`
	Angle    Angle     `json:"angle"`
}

func (p *Projectile) Vector() *Vector {
	return AngleToVector(p.Angle)
}

func (p *Projectile) Tick(t float64) {
	v := p.Vector()
	amount := t * float64(game.constants.ProjectileSpeed)
	p.Position.X += int64(v.X * amount)
	p.Position.Y += int64(v.Y * amount)
}
