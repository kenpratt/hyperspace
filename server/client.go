package main

import (
	"encoding/json"
	"log"
	"time"
)

type Client struct {
	// WebSocket connection (communicate with this via send and receive channels)
	conn *Connection

	// ID of the player that this client represents
	playerId uint16
}

func makeClient(conn *Connection) *Client {
	c := &Client{conn, 0}
	h.register <- c
	return c
}

func (c *Client) Initialize(playerId uint16) {
	c.playerId = playerId

	// send initial player data to client
	b, _ := json.Marshal(PlayerData{c.playerId, 256, 110})
	raw := json.RawMessage(b)
	c.Send(&Message{"init", &raw})

	// boot client message handler
	go c.run()
}

func (c *Client) run() {
	defer func() {
		h.unregister <- c
		close(c.conn.send)
	}()

	for {
		select {
		case message, ok := <-c.conn.receive:
			if !ok {
				log.Println("Client stopping", c.playerId)
				return
			}
			c.handleMessage(message)
		}
	}
}

func (c *Client) handleMessage(message *Message) {
	switch message.Type {
	case "position":
		var pos PositionData
		err := json.Unmarshal([]byte(*message.Data), &pos)
		if err != nil {
			log.Fatal(err)
		}

		b, _ := json.Marshal(PlayerData{c.playerId, pos.X, pos.Y})
		raw := json.RawMessage(b)
		h.broadcast <- &Message{"position", &raw}
	case "fire":
		var data FireData
		err := json.Unmarshal([]byte(*message.Data), &data)
		if err != nil {
			log.Fatal(err)
		}

		// TODO: Why does this close the connection?
		pos := &PositionData{X: 0, Y: 0}
		projectile := ProjectileData{Id: data.Id, Angle: 0, Position: pos}
		go func() {
			for i := 0; i < 1000; i++ {
				projectile.UpdateOneTick()
				b, _ := json.Marshal(projectile)
				raw := json.RawMessage(b)
				h.broadcast <- &Message{"fire", &raw}
				time.Sleep(time.Duration(25) * time.Millisecond)
			}
		}()
	}

}

func (c *Client) Send(message *Message) {
	c.conn.send <- message
}
