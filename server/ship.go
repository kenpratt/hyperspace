package main

type Ship struct {
	Id           string    `json:"id"`
	Position     *Position `json:"position"`
	Angle        Angle     `json:"angle"`
	Acceleration int8      `json:"acceleration"`
	Rotation     int8      `json:"rotation"`
}

func (s *Ship) Tick(t uint64) {
	if s.Rotation != 0 {
		s.Angle += AmountToRotate(s.Rotation, game.constants.ShipRotation, t)
	}

	if s.Acceleration == 1 {
		x, y := AmountToMove(s.Angle, game.constants.ShipAcceleration, t)
		s.Position.X += x
		s.Position.Y += y
	}
}
