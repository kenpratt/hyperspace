package main

type GameState struct {
	// Game objects that currently exist
	Ships       map[string]*Ship       `json:"ships"`
	Projectiles map[string]*Projectile `json:"projectiles"`
	Asteroids   map[string]*Asteroid   `json:"asteroids"`
}

func CreateGameState() *GameState {
	return &GameState{
		Ships:       make(map[string]*Ship),
		Projectiles: make(map[string]*Projectile),
		Asteroids:   make(map[string]*Asteroid),
	}
}

func (s *GameState) Tick(elapsed uint64) *GameState {
	t := CreateGameState()

	for _, o := range s.Ships {
		p := o.Tick(elapsed)
		if p != nil {
			t.Ships[p.Id] = p
		}
	}

	for _, o := range s.Projectiles {
		p := o.Tick(elapsed)
		if p != nil {
			t.Projectiles[p.Id] = p
		}
	}

	for _, o := range s.Asteroids {
		p := o.Tick(elapsed)
		if p != nil {
			t.Asteroids[p.Id] = p
		}
	}

	return t
}
