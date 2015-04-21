package main

import (
	"fmt"
	"log"
	"math"
)

type Ship struct {
	Id           string `json:"id"`
	Position     *Point `json:"position"`
	Angle        Angle  `json:"angle"`
	Acceleration int8   `json:"acceleration"`
	Rotation     int8   `json:"rotation"`
}

const (
	ShipRadius = 10
)

func (s *Ship) Tick(t uint64) {
	if s.Rotation != 0 {
		d := AmountToRotate(s.Rotation, game.constants.ShipRotation, t)
		s.Angle = AddFloatToAngle(s.Angle, d)
	}

	if s.Acceleration == 1 {
		// TODO: When we add drift, move velocity to ship struct, and change it due to acceleration
		velocity := AngleAndSpeedToVector(s.Angle, game.constants.ShipAcceleration)
		x, y := AmountToMove(velocity, t)
		s.Position.X += x
		s.Position.Y += y
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
}
