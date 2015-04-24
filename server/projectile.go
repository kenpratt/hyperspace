package main

type Projectile struct {
	Id       string  `bson:"i"`
	Alive    bool    `bson:"z"`
	Position *Point  `bson:"p"`
	Velocity *Vector `bson:"v"`
	Created  uint64  `bson:"c"`
	Owner    string  `bson:"o"`
	Died     uint64  `bson:"-"`
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
	died := p.Died
	if alive && (t-p.Created) >= 2000 {
		alive = false
		died = t
	}

	// return copy of object with new position and alive status
	return &Projectile{
		Id:       p.Id,
		Alive:    alive,
		Position: pos,
		Velocity: p.Velocity,
		Created:  p.Created,
		Owner:    p.Owner,
		Died:     died,
	}
}
