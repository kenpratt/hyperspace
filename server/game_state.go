package main

type GameState struct {
	Time        uint64                 `json:"time"`
	Ships       map[string]*Ship       `json:"ships"`
	Projectiles map[string]*Projectile `json:"projectiles"`
	Asteroids   map[string]*Asteroid   `json:"asteroids"`
}

func CreateGameState(time uint64) *GameState {
	return &GameState{
		Time:        time,
		Ships:       make(map[string]*Ship),
		Projectiles: make(map[string]*Projectile),
		Asteroids:   make(map[string]*Asteroid),
	}
}

func (s *GameState) Tick(time uint64, elapsed uint64) *GameState {
	t := CreateGameState(time)

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
