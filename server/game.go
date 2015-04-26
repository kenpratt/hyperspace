package main

import (
	"math"
	"strconv"
	"time"
)

type Game struct {
	// Registered clients.
	clients map[*Client]bool

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Game objects that currently exist
	history *GameHistory

	// Next valid game object id.
	nextId int
}

type GameSettings struct {
	// All the important game variables
	constants *GameConstants

	// Whether or not to print debug messages.
	debug bool
}

type GameConstants struct {
	ShipRadius       float64 `json:"ship_radius"`
	ShipAcceleration float64 `json:"ship_acceleration"`
	ShipRotation     float64 `json:"ship_rotation"`
	ShipDrag         float64 `json:"ship_drag"`
	ProjectileRadius float64 `json:"projectile_radius"`
	ProjectileSpeed  float64 `json:"projectile_speed"`
}

type GameError struct {
	What string
}

func (e GameError) Error() string {
	return e.What
}

const (
	gameUpdatePeriod = 10 * time.Millisecond
	updateTimeBuffer = 2000 // in milliseconds
)

// TODO: Get this working without a global variable, I guess pass a ref to game into the web socket handler function?
var settings = &GameSettings{
	debug: false,

	// Game constants, values are all per-second
	constants: &GameConstants{
		ShipRadius:       5.6,  // Pixels
		ShipAcceleration: 100,  // Pixels per second^2
		ShipDrag:         -0.2, // Percentage reduction per second
		ShipRotation:     200,  // Degrees per second
		ProjectileRadius: 3,    // Pixels
		ProjectileSpeed:  180,  // Pixels per second
	},
}
var game = CreateGame()

func CreateGame() *Game {
	g := &Game{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		history:    CreateGameHistory(),
	}

	// Create asteroids
	for i := 0; i < 100; i++ {
		id := g.generateId()
		geom := RandomAsteroidGeometry()
		g.history.Run(&CreateAsteroidEvent{MakeTimestamp(), id, geom})
	}

	return g
}

func (g *Game) Run() {
	gameUpdateTicker := time.NewTicker(gameUpdatePeriod)
	defer func() {
		gameUpdateTicker.Stop()
	}()

	for {
		select {
		case c := <-g.register:
			// Register client
			g.clients[c] = true

			// Create ship
			id := g.generateId()
			color := RandomBrightColor()
			pos := MakePoint(float64(Random(-1000, 1000)), float64(Random(-1000, 1000)))
			state := g.history.Run(&CreateShipEvent{MakeTimestamp(), id, color, pos})

			// Send game state dump to player
			c.Initialize(id, settings.constants, state)
		case c := <-g.unregister:
			if _, ok := g.clients[c]; ok {
				delete(g.clients, c)
			}
		case <-gameUpdateTicker.C:
			g.history.Tick(g.lowestSeenUpdateTime())
			// state := g.history.Tick(g.lowestSeenUpdateTime())
			// if settings.debug {
			// 	log.Println(fmt.Sprintf("Ships: %d, Projectiles: %d", len(state.Ships), len(state.Projectiles)))
			// }
		}
	}
}

func (g *Game) generateId() string {
	g.nextId++
	return strconv.Itoa(g.nextId)
}

func (g *Game) lowestSeenUpdateTime() uint64 {
	var lowest uint64 = math.MaxUint64
	for c, _ := range g.clients {
		t := c.LastUpdateTime() - updateTimeBuffer
		if t < lowest {
			lowest = t
		}
	}
	return lowest
}
