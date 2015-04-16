package main

type Projectile struct {
	Id       string    `json:"id"`
	Position *Position `json:"position"`
	Angle    float64   `json:"angle"`
}

func (p *Projectile) Vector() *Vector {
	return AngleToVector(p.Angle)
}

func (p *Projectile) Tick(t float64) {
	p.Position = &Position{
		X: p.Position.X + p.Vector().X*game.constants.ProjectileSpeed*t,
		Y: p.Position.Y + p.Vector().Y*game.constants.ProjectileSpeed*t,
	}
}
