package main

type Projectile struct {
	Id       string  `json:"i"`
	Alive    bool    `json:"z"`
	Position *Point  `json:"p"`
	Velocity *Vector `json:"v"`
	Created  uint64  `json:"c"`
	Owner    string  `json:"o"`
}

const (
	ProjectileRadius = 10
)

func CreateProjectile(id string, pos *Point, angle float64, created uint64, owner string) *Projectile {

	return &Projectile{
		Id:       id,
		Alive:    true,
		Position: pos,
		Velocity: RoundVector(AngleAndSpeedToVector(angle, settings.constants.ProjectileSpeed)),
		Created:  created,
		Owner:    owner,
	}
}

func (p *Projectile) Tick(t uint64, state *GameState) *Projectile {
	// calculate time since last update (in milliseconds)
	elapsed := t - state.Time

	// calculate new position
	x, y := AmountToMove(p.Velocity, elapsed)
	pos := MakePoint(p.Position.X+x, p.Position.Y+y)

	// calculate new aliveness
	alive := p.Alive
	if alive && (t-p.Created) >= 2000 {
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
