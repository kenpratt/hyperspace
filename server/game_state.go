package main

import "log"

type GameState struct {
	Time        uint64                 `json:"time"`
	Ships       map[string]*Ship       `json:"ships"`
	Projectiles map[string]*Projectile `json:"projectiles"`
	Asteroids   map[string]*Asteroid   `json:"asteroids"`
}

func CreateGameState(t uint64) *GameState {
	return &GameState{
		Time:        t,
		Ships:       make(map[string]*Ship),
		Projectiles: make(map[string]*Projectile),
		Asteroids:   make(map[string]*Asteroid),
	}
}

func (s *GameState) Tick(t uint64) *GameState {
	if t < s.Time {
		log.Fatalf("Tried to call tick with timestamp lower than previous tick: %d, %d", s.Time, t)
		return nil
	}

	// TODO: If t == s.Time, do a clone of game objects and return, since no time will have elapsed?

	// create new state
	state := CreateGameState(t)

	for _, o := range s.Ships {
		p := o.Tick(t, s)
		if p != nil {
			state.Ships[p.Id] = p
		}
	}

	for _, o := range s.Projectiles {
		p := o.Tick(t, s)
		if p != nil {
			state.Projectiles[p.Id] = p
		}
	}

	for _, o := range s.Asteroids {
		p := o.Tick(t, s)
		if p != nil {
			state.Asteroids[p.Id] = p
		}
	}

	return state
}
