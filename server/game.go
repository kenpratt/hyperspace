package main

import (
	"encoding/json"
	"log"
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

	// Next valid game object id.
	nextId uint16
}

type GameError struct {
	What string
}

func (e GameError) Error() string {
	return e.What
}

// TODO get this working without a global variable, I guess pass a ref to game into the web socket handler function?
var game = Game{
	clients:    make(map[*Client]bool),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	events:     make(chan *Event),
}

func (g *Game) run() {
	for {
		select {
		case c := <-g.register:
			g.clients[c] = true
			g.nextId++
			c.Initialize(g.nextId)
		case c := <-g.unregister:
			if _, ok := g.clients[c]; ok {
				delete(g.clients, c)
			}
		case e := <-g.events:
			updates, err := g.applyEvent(e)
			if err != nil {
				log.Println("Error applying event", e, err)
				continue
			}

			for c := range g.clients {
				c.Send(updates)
			}
		}
	}
}

func (g *Game) applyEvent(e *Event) (*Message, error) {
	switch e.Type {
	case "position":
		data := e.Data.(*PlayerData)
		b, _ := json.Marshal(*data)
		raw := json.RawMessage(b)
		return &Message{"position", &raw}, nil
	case "fire":
		data := e.Data.(*ProjectileData)
		b, _ := json.Marshal(*data)
		raw := json.RawMessage(b)
		return &Message{"position", &raw}, nil
	default:
		return nil, GameError{"Don't know how to apply event"}
	}

}
