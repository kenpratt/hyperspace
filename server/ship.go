package main

import (
	"log"
)

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

func (s *Ship) Tick() {
	if s.Rotation != 0 {
		s.Angle += s.Rotation * 2
	}

	if s.Acceleration == 1 {
		log.Println("accelerating", s.Angle, s.Vector(), s.Position)
		s.Position = &Position{
			X: s.Position.X + s.Vector().X*5,
			Y: s.Position.Y + s.Vector().Y*5,
		}
	}
}
