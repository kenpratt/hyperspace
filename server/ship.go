package main

type Ship struct {
	Id           string    `json:"id"`
	Position     *Position `json:"position"`
	Angle        Angle     `json:"angle"`
	Acceleration int8      `json:"acceleration"`
	Rotation     int8      `json:"rotation"`
}

func (s *Ship) Vector() *Vector {
	return AngleToVector(s.Angle)
}

func (s *Ship) Tick(t float64) {
	amount := t * float64(game.constants.ShipRotation)
	if s.Rotation != 0 {
		s.Angle += Angle(float64(s.Rotation) * amount)
	}

	v := s.Vector()
	amount = t * float64(game.constants.ProjectileSpeed)
	if s.Acceleration == 1 {
		s.Position.X += int64(v.X * amount)
		s.Position.Y += int64(v.Y * amount)
	}
}
