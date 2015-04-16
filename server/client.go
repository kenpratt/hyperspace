package main

import (
	"encoding/json"
	"log"
)

// A Client is a connected player and associated websocket connection.
type Client struct {
	// WebSocket connection (communicate with this via send and receive channels)
	conn *Connection

	// ID of the player that this client represents
	playerId string
}

func makeClient(conn *Connection) *Client {
	c := &Client{conn: conn}
	game.register <- c
	return c
}

func (c *Client) Initialize(playerId string, gameConstants *GameConstants, gameState *UpdateData) {
	c.playerId = playerId

	// send initial player data to client
	b, err := json.Marshal(&InitData{playerId, gameConstants, gameState})
	if err != nil {
		panic(err)
	}
	raw := json.RawMessage(b)
	c.Send(&Message{"init", MakeTimestamp(), &raw})

	// boot client message handler
	go c.run()
}

func (c *Client) run() {
	defer func() {
		game.unregister <- c
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
	case "changeAcceleration":
		var data AccelerationData
		err := json.Unmarshal([]byte(*message.Data), &data)
		if err != nil {
			log.Fatal(err)
		}

		game.events <- &ChangeAccelerationEvent{c.playerId, message.Time, data.Direction}
	case "changeRotation":
		var data RotationData
		err := json.Unmarshal([]byte(*message.Data), &data)
		if err != nil {
			log.Fatal(err)
		}

		game.events <- &ChangeRotationEvent{c.playerId, message.Time, data.Direction}
	case "fire":
		var data FireData
		err := json.Unmarshal([]byte(*message.Data), &data)
		if err != nil {
			log.Fatal(err)
		}

		game.events <- &FireEvent{c.playerId, message.Time, data.ProjectileId}
	}

}

func (c *Client) Send(message *Message) {
	c.conn.send <- message
}
