package main

type Projectile struct {
	Id       string    `json:"id"`
	Position *Position `json:"position"`
	Angle    float64   `json:"angle"`
}

func (p *Projectile) Vector() *Vector {
	return AngleToVector(p.Angle)
}

func (p *Projectile) UpdateOneTick() {
	p.Position = &Position{
		X: p.Position.X + p.Vector().X,
		Y: p.Position.Y + p.Vector().Y,
	}
}
