package main

type GameEvent interface {
	Execute(*GameState) error
}

// create a ship
type CreateShipEvent struct {
	Time     uint64
	Id       string
	Position *Point
}

func (e *CreateShipEvent) Execute(state *GameState) error {
	state.Ships[e.Id] = CreateShip(e.Id, e.Position)
	return nil
}

// create an asteroid
type CreateAsteroidEvent struct {
	Time     uint64
	Id       string
	Position *Point
	Angle    float64
	Velocity *Vector
	Shape    []*Point
}

func (e *CreateAsteroidEvent) Execute(state *GameState) error {
	state.Asteroids[e.Id] = CreateAsteroid(e.Id, e.Position, e.Angle, e.Velocity, e.Shape)
	return nil
}

// remove a ship
type RemoveShipEvent struct {
	Time   uint64
	ShipId string
}

func (e *RemoveShipEvent) Execute(state *GameState) error {
	delete(state.Ships, e.ShipId)
	return nil
}

// change ship acceleration direction
type ChangeAccelerationEvent struct {
	Time      uint64
	ShipId    string
	Direction int8
}

func (e *ChangeAccelerationEvent) Execute(state *GameState) error {
	s := state.Ships[e.ShipId]
	if s == nil {
		return GameError{"Ship doesn't exist for player"}
	}

	s.Acceleration = e.Direction
	return nil
}

// change ship rotation direction
type ChangeRotationEvent struct {
	Time      uint64
	ShipId    string
	Direction int8
}

func (e *ChangeRotationEvent) Execute(state *GameState) error {
	s := state.Ships[e.ShipId]
	if s == nil {
		return GameError{"Ship doesn't exist for player"}
	}

	s.Rotation = e.Direction
	return nil
}

// fire ship laser!
type FireEvent struct {
	Time         uint64
	ShipId       string
	ProjectileId string
	Created      uint64
}

func (e *FireEvent) Execute(state *GameState) error {
	s := state.Ships[e.ShipId]
	if s == nil {
		return GameError{"Ship doesn't exist for player"}
	}

	pos := *s.Position // Clone ship position
	projectile := CreateProjectile(e.ProjectileId, &pos, s.Angle, e.Created, e.ShipId)
	state.Projectiles[projectile.Id] = projectile
	return nil
}

// remove dead objects
type CleanupEvent struct {
	Time uint64
}

func (e *CleanupEvent) Execute(state *GameState) error {
	dead := []string{}

	for k, v := range state.Projectiles {
		if !v.Alive {
			dead = append(dead, k)
		}
	}

	for i := range dead {
		delete(state.Projectiles, dead[i])
	}

	return nil
}
