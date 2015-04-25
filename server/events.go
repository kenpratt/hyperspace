package main

type GameEvent interface {
	Time() uint64
	Execute(*GameState) error
}

// create a ship
type CreateShipEvent struct {
	time     uint64
	id       string
	position *Point
}

func (e *CreateShipEvent) Time() uint64 {
	return e.time
}

func (e *CreateShipEvent) Execute(state *GameState) error {
	state.Ships[e.id] = CreateShip(e.id, e.position)
	return nil
}

// create an asteroid
type CreateAsteroidEvent struct {
	time     uint64
	id       string
	position *Point
	angle    float64
	velocity *Vector
	shape    []*Point
}

func (e *CreateAsteroidEvent) Time() uint64 {
	return e.time
}

func (e *CreateAsteroidEvent) Execute(state *GameState) error {
	state.Asteroids[e.id] = CreateAsteroid(e.id, e.position, e.angle, e.velocity, e.shape)
	return nil
}

// remove a ship
type RemoveShipEvent struct {
	time   uint64
	shipId string
}

func (e *RemoveShipEvent) Time() uint64 {
	return e.time
}

func (e *RemoveShipEvent) Execute(state *GameState) error {
	delete(state.Ships, e.shipId)
	return nil
}

// change ship acceleration direction
type ChangeAccelerationEvent struct {
	time      uint64
	shipId    string
	direction int8
}

func (e *ChangeAccelerationEvent) Time() uint64 {
	return e.time
}

func (e *ChangeAccelerationEvent) Execute(state *GameState) error {
	s := state.Ships[e.shipId]
	if s == nil {
		return GameError{"Ship doesn't exist for player"}
	}

	s.Acceleration = e.direction
	return nil
}

// change ship rotation direction
type ChangeRotationEvent struct {
	time      uint64
	shipId    string
	direction int8
}

func (e *ChangeRotationEvent) Time() uint64 {
	return e.time
}

func (e *ChangeRotationEvent) Execute(state *GameState) error {
	s := state.Ships[e.shipId]
	if s == nil {
		return GameError{"Ship doesn't exist for player"}
	}

	s.Rotation = e.direction
	return nil
}

// fire ship laser!
type FireEvent struct {
	time         uint64
	shipId       string
	projectileId string
	created      uint64
}

func (e *FireEvent) Time() uint64 {
	return e.time
}

func (e *FireEvent) Execute(state *GameState) error {
	s := state.Ships[e.shipId]
	if s == nil {
		return GameError{"Ship doesn't exist for player"}
	}

	pos := *s.Position // Clone ship position
	projectile := CreateProjectile(e.projectileId, &pos, s.Angle, s.Velocity, e.created, e.shipId)
	state.Projectiles[projectile.Id] = projectile
	return nil
}

// game tick event
type TickEvent struct {
	time        uint64
	syncedUntil uint64
}

func (e *TickEvent) Time() uint64 {
	return e.time
}

func (e *TickEvent) Execute(state *GameState) error {
	// delete objects that died long enough ago that all clients have seen their dead status
	for id, p := range state.Ships {
		if !p.Alive && p.Died <= e.syncedUntil {
			delete(state.Ships, id)
		}
	}
	for id, p := range state.Projectiles {
		if !p.Alive && p.Died <= e.syncedUntil {
			delete(state.Projectiles, id)
		}
	}
	for id, p := range state.Asteroids {
		if !p.Alive && p.Died <= e.syncedUntil {
			delete(state.Asteroids, id)
		}
	}
	return nil
}
