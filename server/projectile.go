package main

type Projectile struct {
	Id       string  `json:"i"`
	Alive    bool    `json:"z"`
	Died     uint64  `json:"-"`
	Position *Point  `json:"p"`
	Velocity *Vector `json:"v"`
	Created  uint64  `json:"c"`
	Owner    string  `json:"o"`
}

func CreateProjectile(id string, pos *Point, angle float64, velocity *Vector, created uint64, owner string) *Projectile {

	return &Projectile{
		Id:       id,
		Alive:    true,
		Died:     0,
		Position: pos,
		Velocity: RoundVector(AddVectors(velocity, AngleAndSpeedToVector(angle, settings.constants.ProjectileSpeed))),
		Created:  created,
		Owner:    owner,
	}
}

func (p *Projectile) Tick(t uint64, state *GameState) *Projectile {
	// calculate time since last update (in milliseconds)
	elapsedMillis := t - state.Time

	// elapsed time in percentage of a second
	elapsed := float64(elapsedMillis) / 1000

	// calculate new position
	pos := MakePoint(p.Position.X+p.Velocity.X*elapsed, p.Position.Y+p.Velocity.Y*elapsed)

	// calculate new aliveness
	alive := p.Alive
	died := p.Died
	if alive && (t-p.Created) >= 2000 {
		alive = false
		died = t
	}

	// return copy of object with new position and alive status
	return &Projectile{
		Id:       p.Id,
		Alive:    alive,
		Died:     died,
		Position: pos,
		Velocity: p.Velocity,
		Created:  p.Created,
		Owner:    p.Owner,
	}
}
