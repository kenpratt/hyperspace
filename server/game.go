package main

import (
	"encoding/json"
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

	// Inbound events from the clients/AIs.
	events chan *Event

	// Game objects that currently exist
	ships       map[string]*Ship
	projectiles map[string]*Projectile

	// Next valid game object id.
	nextId int
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
	events:      make(chan *Event),
	ships:       make(map[string]*Ship),
	projectiles: make(map[string]*Projectile),
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
			c.Initialize(id, g.fullState())
		case c := <-g.unregister:
			if _, ok := g.clients[c]; ok {
				delete(g.clients, c)
			}

			// TODO mark the ship as dead, and then after a time, delete them from the list of ships
			// delete(g.ships, c.shipId)
		case e := <-g.events:
			err := g.applyEvent(e)
			if err != nil {
				log.Println("Error applying event", e, err)
				continue
			}
		case <-ticker.C:
			g.broadcastUpdate()
		}
	}
}

func (g *Game) applyEvent(e *Event) error {
	s := g.ships[e.PlayerId]
	if s == nil {
		return GameError{"Ship doesn't exist for player"}
	}

	switch e.Type {
	case "position":
		// TODO change movement to be based on start/stop/rotate instead of position updates
		data := e.Data.(*PositionData)

		// move the player
		s.Position.X = data.X
		s.Position.Y = data.X
		return nil
	case "fire":
		data := e.Data.(*FireData)

		// create a projectile and spawn goroutine to move it forward (TODO switch to game loop)
		pos := &Position{X: s.Position.X, Y: s.Position.Y}
		projectile := Projectile{Id: data.Id, Angle: 0, Position: pos}
		g.projectiles[projectile.Id] = &projectile
		go func() {
			// TODO get rid of this goroutine, and move logic into a game loop that updates all physics at the same time
			// TODO handle projectile death in a nicer way
			for i := 0; i < 1000; i++ {
				projectile.UpdateOneTick()
				g.broadcastUpdate()
				time.Sleep(time.Duration(25) * time.Millisecond)
			}
		}()
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
	m := &Message{"update", &raw}

	for c := range g.clients {
		c.Send(m)
	}
}
