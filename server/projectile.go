package main

type Projectile struct {
	Id       string  `json:"id"`
	Alive    bool    `json:"alive"`
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
		Alive:    true,
		Position: pos,
		Velocity: AngleAndSpeedToVector(angle, game.constants.ProjectileSpeed),
		Created:  created,
		Owner:    owner,
	}
}

func (p *Projectile) Tick(elapsed uint64, state *GameState) *Projectile {
	// calculate new position
	x, y := AmountToMove(p.Velocity, elapsed)
	pos := &Point{p.Position.X + x, p.Position.Y + y}

	// calculate new aliveness
	alive := p.Alive
	if alive && (MakeTimestamp()-p.Created) >= 2000 {
		alive = false
	}

	// return copy of object with new position
	return &Projectile{
		Id:       p.Id,
		Alive:    alive,
		Position: pos,
		Velocity: p.Velocity,
		Created:  p.Created,
		Owner:    p.Owner,
	}
}
