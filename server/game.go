package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type Game struct {
	// All the important game variables
	constants *GameConstants

	// Registered clients.
	clients map[*Client]bool

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Inbound events from the clients/AIs.
	events chan interface{}

	// Game objects that currently exist
	ships       map[string]*Ship
	projectiles map[string]*Projectile

	// Next valid game object id.
	nextId int
}

type GameConstants struct {
	ShipAcceleration uint16 `json:"ship_acceleration"`
	ShipRotation     uint16 `json:"ship_rotation"`
	ProjectileSpeed  uint16 `json:"projectile_speed"`
}

type GameError struct {
	What string
}

func (e GameError) Error() string {
	return e.What
}

const (
	updatePeriod = 100 * time.Millisecond
)

// TODO: Get this working without a global variable, I guess pass a ref to game into the web socket handler function?
var game = Game{
	clients:     make(map[*Client]bool),
	register:    make(chan *Client),
	unregister:  make(chan *Client),
	events:      make(chan interface{}),
	ships:       make(map[string]*Ship),
	projectiles: make(map[string]*Projectile),

	// Game constants, values per second
	constants: &GameConstants{
		ShipAcceleration: 100, // Pixels per second
		ShipRotation:     100, // Degrees per second
		ProjectileSpeed:  150, // Pixels per second
	},
}

func (g *Game) run() {
	ticker := time.NewTicker(updatePeriod)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case c := <-g.register:
			// Register client
			g.clients[c] = true

			// Create ship
			g.nextId++
			id := strconv.Itoa(g.nextId)
			pos := &Position{X: 256, Y: 110}
			s := &Ship{Id: id, Angle: 0, Position: pos}
			g.ships[id] = s

			// Send game state dump to player
			c.Initialize(id, g.constants, g.fullState())
		case c := <-g.unregister:
			if _, ok := g.clients[c]; ok {
				delete(g.clients, c)
			}

			// TODO: Mark the ship as dead, and then after a time, delete them from the list of ships.
			// delete(g.ships, c.shipId)
		case e := <-g.events:
			err := g.applyEvent(e)
			if err != nil {
				log.Println("Error applying event", e, err)
				continue
			}
		case <-ticker.C:
			// TODO: Calculate actual elapsed time.
			var elapsed uint64 = 100 // (milliseconds)
			for _, o := range g.ships {
				o.Tick(elapsed)
			}
			for _, o := range g.projectiles {
				o.Tick(elapsed)
			}
			g.broadcastUpdate()
		}
	}
}

func (g *Game) applyEvent(o interface{}) error {
	switch e := o.(type) {
	case *ChangeAccelerationEvent:
		s := g.ships[e.PlayerId]
		if s == nil {
			return GameError{"Ship doesn't exist for player"}
		}

		s.Acceleration = e.Direction
		return nil
	case *ChangeRotationEvent:
		s := g.ships[e.PlayerId]
		if s == nil {
			return GameError{"Ship doesn't exist for player"}
		}

		s.Rotation = e.Direction
		return nil
	case *FireEvent:
		s := g.ships[e.PlayerId]
		if s == nil {
			return GameError{"Ship doesn't exist for player"}
		}

		pos := *s.Position // Clone ship position
		projectile := Projectile{Id: e.ProjectileId, Angle: s.Angle, Position: &pos}
		g.projectiles[projectile.Id] = &projectile
		return nil
	default:
		return GameError{"Don't know how to apply event"}
	}

}

func (g *Game) fullState() *UpdateData {
	return &UpdateData{g.ships, g.projectiles}
}

func (g *Game) broadcastUpdate() {
	data := g.fullState()
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	raw := json.RawMessage(b)
	m := &Message{"update", MakeTimestamp(), &raw}

	for c := range g.clients {
		c.Send(m)
	}
}
