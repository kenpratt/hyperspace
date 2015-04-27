package main

import (
	"fmt"
	"log"
)

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
		log.Printf("Tried to call tick with timestamp lower than previous tick: %d, %d", s.Time, t)
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

	// TODO: Come up with a better way to look up collisions.
	// From https://developer.mozilla.org/en-US/docs/Games/Techniques/2D_collision_detection

	for _, s := range state.Ships {
		if !s.Alive {
			continue
		}

		// check for ship-ship collisions
		for _, os := range state.Ships {
			if !s.Alive || !os.Alive || s.Id == os.Id {
				continue
			}

			if IsColliding(s.Position, settings.constants.ShipRadius, os.Position, settings.constants.ShipRadius) {
				log.Println(fmt.Sprintf("Ship %v collided with Ship %v", s.Id, os.Id))
				s.Alive = false
				s.Died = t
				os.Alive = false
				os.Died = t
			}
		}

		if !s.Alive {
			continue
		}

		// check for ship-projectile collisions
		for _, p := range state.Projectiles {
			if !s.Alive || !p.Alive || s.Id == p.Owner {
				continue
			}

			if IsColliding(s.Position, settings.constants.ShipRadius, p.Position, settings.constants.ProjectileRadius) {
				log.Println(fmt.Sprintf("Ship %v collided with Projectile %v", s.Id, p.Id))
				s.Alive = false
				s.Died = t
				p.Alive = false
				p.Died = t
			}
		}

		if !s.Alive {
			continue
		}

		// check for ship-asteroid collisions
		for _, a := range state.Asteroids {
			if !s.Alive || !a.Alive {
				continue
			}

			if IsColliding(s.Position, settings.constants.ShipRadius, a.Position, a.Radius) {
				log.Println(fmt.Sprintf("Ship %v collided with Asteroid %v", s.Id, a.Id))
				s.Alive = false
				s.Died = t
			}
		}
	}

	for _, a := range state.Asteroids {
		if !a.Alive {
			continue
		}

		// check for asteroid-projectile collisions
		for _, p := range state.Projectiles {
			if !a.Alive || !p.Alive {
				continue
			}

			if IsColliding(a.Position, a.Radius, p.Position, settings.constants.ProjectileRadius) {
				log.Println(fmt.Sprintf("Asteroid %v collided with Projectile %v", a.Id, p.Id))

				// create new asteroids
				if a.Splittable() {
					a1, a2 := a.Split()
					state.Asteroids[a1.Id] = a1
					state.Asteroids[a2.Id] = a2
				}

				// destroy asteroid and projectile
				a.Alive = false
				a.Died = t
				p.Alive = false
				p.Died = t

			}
		}

		// TODO: Asteroids should bounce off of each other
		// check for asteroid-asteroid collisions
		// for _, oa := range state.Asteroids {
		// 	if !a.Alive || !oa.Alive || a.Id == oa.Id {
		// 		continue
		// 	}
		//
		// 	if IsColliding(a.Position, a.Radius, oa.Position, oa.Radius) {
		// 		log.Println(fmt.Sprintf("Asteroid %v collided with Asteroid %v", a.Id, oa.Id))
		// 	}
		// }
	}

	return state
}
