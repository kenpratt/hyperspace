package main

import (
	"encoding/json"
	"fmt"
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
	asteroids   map[string]*Asteroid

	// Next valid game object id.
	nextId int

	// Whether or not to print debug messages.
	debug bool

	// Last physics update time
	lastUpdate uint64
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
	clientUpdatePeriod = 100 * time.Millisecond
	gameTickPeriod     = 10 * time.Millisecond
)

// TODO: Get this working without a global variable, I guess pass a ref to game into the web socket handler function?
var game = CreateGame()

func CreateGame() *Game {
	g := &Game{
		clients:     make(map[*Client]bool),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		events:      make(chan interface{}),
		ships:       make(map[string]*Ship),
		projectiles: make(map[string]*Projectile),
		asteroids:   make(map[string]*Asteroid),
		debug:       false,

		// Game constants, values per second
		constants: &GameConstants{
			ShipAcceleration: 100, // Pixels per second
			ShipRotation:     100, // Degrees per second
			ProjectileSpeed:  150, // Pixels per second
		},

		// beginning of game time
		lastUpdate: MakeTimestamp(),
	}

	// Generate asteroids
	for i := 0; i < 100; i++ {
		id := g.generateId()
		g.asteroids[id] = CreateAsteroid(id)
	}

	return g
}

func (g *Game) run(debug bool) {
	g.debug = debug
	updateTicker := time.NewTicker(gameTickPeriod)
	gameTicker := time.NewTicker(clientUpdatePeriod)
	defer func() {
		updateTicker.Stop()
		gameTicker.Stop()
	}()

	for {
		select {
		case c := <-g.register:
			// Register client
			g.clients[c] = true

			// Create ship
			id := g.generateId()
			pos := &Coordinate{X: 0, Y: 0}
			s := &Ship{Id: id, Angle: 0, Position: pos}
			g.ships[id] = s

			// Send game state dump to player
			c.Initialize(id, g.constants, g.fullState())
		case c := <-g.unregister:
			if _, ok := g.clients[c]; ok {
				delete(g.clients, c)
			}
		case e := <-g.events:
			err := g.applyEvent(e)
			if err != nil {
				log.Println("Error applying event", e, err)
				continue
			}
		case <-gameTicker.C:
			err := g.cleanup()
			if err != nil {
				log.Println("Error Cleaning Up", err)
				continue
			}

			// calculate time since last update (in milliseconds)
			now := MakeTimestamp()
			g.lastUpdate = now
			elapsed := uint64(gameTickPeriod / time.Millisecond)

			// update physics
			for _, o := range g.ships {
				o.Tick(elapsed)
			}
			for _, o := range g.projectiles {
				o.Tick(elapsed)
			}
			for _, o := range g.asteroids {
				o.Tick(elapsed)
			}
		case <-updateTicker.C:
			g.broadcastUpdate(g.lastUpdate)
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
		projectile := Projectile{
			Id:       e.ProjectileId,
			Velocity: AngleAndSpeedToVector(s.Angle, game.constants.ProjectileSpeed),
			Position: &pos,
			Created:  e.Created,
			Owner:    e.PlayerId,
		}
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
			dead = append(dead, k)
		}
	}

	for i := range dead {
		delete(g.projectiles, dead[i])
	}

	return nil
}

func (g *Game) fullState() *UpdateData {
	return &UpdateData{g.ships, g.projectiles, g.asteroids}
}

func (g *Game) broadcastUpdate(t uint64) {
	data := g.fullState()
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	raw := json.RawMessage(b)
	m := &Message{"update", t, &raw}

	if g.debug {
		log.Println(fmt.Sprintf("Ships: %d, Projectiles: %d", len(g.ships), len(g.projectiles)))
	}

	for c := range g.clients {
		c.Send(m)
	}
}

func (g *Game) generateId() string {
	g.nextId++
	return strconv.Itoa(g.nextId)
}
