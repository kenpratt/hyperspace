package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type Game struct {
	// all the important game variables
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

// TODO get this working without a global variable, I guess pass a ref to game into the web socket handler function?
var game = Game{
	clients:     make(map[*Client]bool),
	register:    make(chan *Client),
	unregister:  make(chan *Client),
	events:      make(chan interface{}),
	ships:       make(map[string]*Ship),
	projectiles: make(map[string]*Projectile),

	// game constants, values per second
	constants: &GameConstants{
		ShipAcceleration: 100, // pixels per second
		ShipRotation:     100, // degrees per second
		ProjectileSpeed:  150, // pixels per second
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
			// register client
			g.clients[c] = true

			// create ship
			g.nextId++
			id := strconv.Itoa(g.nextId)
			pos := &Position{X: 256, Y: 110}
			s := &Ship{Id: id, Angle: 0, Position: pos}
			g.ships[id] = s

			// send game state dump to player
			c.Initialize(id, g.constants, g.fullState())
		case c := <-g.unregister:
			if _, ok := g.clients[c]; ok {
				delete(g.clients, c)
			}

			// TODO: mark the ship as dead, and then after a time, delete them from the list of ships
			// delete(g.ships, c.shipId)
		case e := <-g.events:
			err := g.applyEvent(e)
			if err != nil {
				log.Println("Error applying event", e, err)
				continue
			}
		case <-ticker.C:
			err := g.cleanup()
			if err != nil {
				log.Println("Error Cleaning Up", err)
				continue
			}

			// TODO: calculate real elapsed time
			elapsed := 0.1 // in seconds
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

		// create a projectile and spawn goroutine to move it forward (TODO switch to game loop)
		pos := &Position{X: s.Position.X, Y: s.Position.Y}
		projectile := Projectile{Id: e.ProjectileId, Angle: s.Angle, Position: pos, Created: e.Created}
		g.projectiles[projectile.Id] = &projectile
		return nil
	default:
		return GameError{"Don't know how to apply event"}
	}
}

func (g *Game) cleanup() error {
	dead := []string{}

	for k, v := range g.projectiles {
		if !v.Alive() {
			log.Println("killing", k, v)
			dead = append(dead, k)
		}
	}

	for i := range dead {
		delete(g.projectiles, dead[i])
	}

	return nil
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

	// log.Println(fmt.Sprintf("Ships: %d, Projectiles: %d", len(g.ships), len(g.projectiles)))

	for c := range g.clients {
		c.Send(m)
	}
}
