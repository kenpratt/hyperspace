package main

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
