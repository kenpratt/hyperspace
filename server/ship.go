package main

type Ship struct {
	Id           string  `json:"i"`
	Color        string  `json:"k"`
	Alive        bool    `json:"z"`
	Died         uint64  `json:"-"`
	Position     *Point  `json:"p"`
	Angle        float64 `json:"a"`
	Velocity     *Vector `json:"v"`
	Acceleration int8    `json:"l"`
	Rotation     int8    `json:"r"`
}

func CreateShip(id string, color string, pos *Point) *Ship {
	return &Ship{
		Id:           id,
		Color:        color,
		Alive:        true,
		Died:         0,
		Position:     pos,
		Angle:        0,
		Velocity:     &Vector{0, 0},
		Acceleration: 0,
		Rotation:     0,
	}
}

func (s *Ship) Tick(t uint64, state *GameState) *Ship {
	// Calculate time since last update (in milliseconds)
	elapsedMillis := t - state.Time

	// Elapsed time in percentage of a second
	elapsed := float64(elapsedMillis) / 1000

	// Calculate new angle
	angle := s.Angle
	if s.Rotation != 0 {
		angle = s.Angle + (settings.constants.ShipRotation * elapsed * float64(s.Rotation))
		for angle < 0 {
			angle += 360
		}
		for angle >= 360 {
			angle -= 360
		}
		angle = RoundToPlaces(angle, 1)
	}

	// Calculate new velocity
	vel := s.Velocity
	if s.Acceleration == 1 {
		accel := AngleAndSpeedToVector(angle, settings.constants.ShipAcceleration)
		vel = AddVectors(s.Velocity, MultiplyVector(accel, elapsed))
	}

	// Apply drag
	vel = AddVectors(vel, MultiplyVector(s.Velocity, settings.constants.ShipDrag*elapsed))

	// Calculate new position
	pos := MakePoint(s.Position.X+vel.X*elapsed, s.Position.Y+vel.Y*elapsed)

	return &Ship{
		Id:           s.Id,
		Color:        s.Color,
		Alive:        s.Alive,
		Died:         s.Died,
		Position:     pos,
		Angle:        angle,
		Velocity:     vel,
		Acceleration: s.Acceleration,
		Rotation:     s.Rotation,
	}
}
