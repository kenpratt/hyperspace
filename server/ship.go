package main

import (
	"fmt"
	"log"
	"math"
)

type Ship struct {
	Id           string  `json:"id"`
	Position     *Point  `json:"position"`
	Angle        float64 `json:"angle"`
	Acceleration int8    `json:"acceleration"`
	Rotation     int8    `json:"rotation"`
}

const (
	ShipRadius = 10
)

func CreateShip(id string, pos *Point) *Ship {
	return &Ship{
		Id:           id,
		Position:     pos,
		Angle:        0,
		Acceleration: 0,
		Rotation:     0,
	}
}

func (s *Ship) Tick(t uint64) *Ship {
	// calculate new angle
	angle := s.Angle
	if s.Rotation != 0 {
		angle = s.Angle + AmountToRotate(s.Rotation, game.constants.ShipRotation, t)
	}

	// calculate new position
	pos := s.Position
	if s.Acceleration == 1 {
		// TODO: When we add drift, move velocity to ship struct, and change it due to acceleration
		velocity := AngleAndSpeedToVector(s.Angle, game.constants.ShipAcceleration)
		x, y := AmountToMove(velocity, t)
		pos = &Point{s.Position.X + x, s.Position.Y + y}
	}

	// TODO: Come up with a better way to look up collisions.
	// From https://developer.mozilla.org/en-US/docs/Games/Techniques/2D_collision_detection
	for _, os := range game.ships {
		if os.Id != s.Id {
			dx := s.Position.X - os.Position.X
			dy := s.Position.Y - os.Position.Y
			distance := math.Sqrt(float64(dx*dx + dy*dy))

			if distance < ShipRadius*2 {
				if game.debug {
					log.Println(fmt.Sprintf("Ship %v colliding with Ship %v", s.Id, os.Id))
				}
			}
		}
	}

	for _, p := range game.projectiles {
		dx := s.Position.X - p.Position.X
		dy := s.Position.Y - p.Position.Y
		distance := math.Sqrt(float64(dx*dx + dy*dy))

		if distance < ShipRadius+ProjectileRadius {
			if game.debug {
				log.Println(fmt.Sprintf("Ship %v colliding with Projectile %v", s.Id, p.Id))
			}
		}
	}

	return &Ship{
		Id:           s.Id,
		Position:     pos,
		Angle:        angle,
		Acceleration: s.Acceleration,
		Rotation:     s.Rotation,
	}
}
