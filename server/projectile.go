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

func CreateProjectile(id string, pos *Point, angle float64, created uint64, owner string) *Projectile {

	return &Projectile{
		Id:       id,
		Position: pos,
		Velocity: AngleAndSpeedToVector(angle, game.constants.ProjectileSpeed),
		Created:  created,
		Owner:    owner,
	}
}

func (p *Projectile) Tick(t uint64) *Projectile {
	// calculate new position
	x, y := AmountToMove(p.Velocity, t)
	pos := &Point{p.Position.X + x, p.Position.Y + y}

	// return copy of object with new position
	return &Projectile{
		Id:       p.Id,
		Position: pos,
		Velocity: p.Velocity,
		Created:  p.Created,
		Owner:    p.Owner,
	}
}

func (p *Projectile) Alive() bool {
	return (MakeTimestamp() - p.Created) < 2000
}
