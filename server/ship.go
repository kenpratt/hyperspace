package main

type Ship struct {
	Id           string    `json:"id"`
	Position     *Position `json:"position"`
	Angle        float64   `json:"angle"`
	Acceleration float64   `json:"acceleration"`
	Rotation     float64   `json:"rotation"`
}

func (s *Ship) Vector() *Vector {
	return AngleToVector(s.Angle)
}

func (s *Ship) Tick(t float64) {
	if s.Rotation != 0 {
		s.Angle += s.Rotation * game.constants.ShipRotation * t
	}

	if s.Acceleration == 1 {
		s.Position = &Position{
			X: s.Position.X + s.Vector().X*game.constants.ShipAcceleration*t,
			Y: s.Position.Y + s.Vector().Y*game.constants.ShipAcceleration*t,
		}
	}
}
