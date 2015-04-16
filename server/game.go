package main

type Game struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Next valid game object id.
	nextId uint16
}

// TODO get this working without a global variable, I guess pass a ref to game into the web socket handler function?
var game = Game{
	broadcast:  make(chan *Message),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
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
		case m := <-g.broadcast:
			for c := range g.clients {
				c.Send(m)
			}
		}
	}
}
