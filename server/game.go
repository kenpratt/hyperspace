package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	gameUpdatePeriod   = 10 * time.Millisecond
	clientUpdatePeriod = 100 * time.Millisecond
)

// TODO: Get this working without a global variable, I guess pass a ref to game into the web socket handler function?
var settings = &GameSettings{
	debug: false,

	// Game constants, values are all per-second
	constants: &GameConstants{
		ShipAcceleration: 100, // Pixels per second
		ShipRotation:     100, // Degrees per second
		ProjectileSpeed:  150, // Pixels per second
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
		p, a, v, s := RandomAsteroidGeometry()
		g.history.Run(&CreateAsteroidEvent{MakeTimestamp(), id, p, a, v, s})
	}

	return g
}

func (g *Game) Run() {
	gameUpdateTicker := time.NewTicker(gameUpdatePeriod)
	clientUpdateTicker := time.NewTicker(clientUpdatePeriod)
	defer func() {
		gameUpdateTicker.Stop()
		clientUpdateTicker.Stop()
	}()

	for {
		select {
		case c := <-g.register:
			// Register client
			g.clients[c] = true

			// Create ship
			id := g.generateId()
			state := g.history.Run(&CreateShipEvent{MakeTimestamp(), id, &Point{X: 0, Y: 0}})

			// Send game state dump to player
			c.Initialize(id, settings.constants, state)
		case c := <-g.unregister:
			if _, ok := g.clients[c]; ok {
				delete(g.clients, c)
			}
		case <-gameUpdateTicker.C:
			g.history.Tick()
		case <-clientUpdateTicker.C:
			state := g.history.GetCurrentStateAndClean()
			g.broadcastUpdate(state)
		}
	}
}

func (g *Game) broadcastUpdate(state *GameState) {
	b, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}

	raw := json.RawMessage(b)
	m := &Message{"update", MakeTimestamp(), &raw}

	if settings.debug {
		log.Println(fmt.Sprintf("Ships: %d, Projectiles: %d", len(state.Ships), len(state.Projectiles)))
	}

	for c := range g.clients {
		c.Send(m)
	}
}

func (g *Game) generateId() string {
	g.nextId++
	return strconv.Itoa(g.nextId)
}
