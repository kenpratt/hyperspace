package main

import (
	"encoding/json"
	"fmt"
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

func (c *Client) Initialize(playerId string, gameConstants *GameConstants, gameState *GameState) {
	c.playerId = playerId

	// send initial player data to client
	b, err := json.Marshal(&InitData{playerId, gameConstants, gameState})
	if err != nil {
		panic(err)
	}
	raw := json.RawMessage(b)
	c.Send(&Message{"init", MakeTimestamp(), &raw})

	log.Println(fmt.Sprintf("Client Starting: %v", c.playerId))

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
				log.Println(fmt.Sprintf("Client Stopping: %v", c.playerId))
				game.history.Run(&RemoveShipEvent{MakeTimestamp(), c.playerId})
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
		game.history.Run(&ChangeAccelerationEvent{message.Time, c.playerId, data.Direction})
	case "changeRotation":
		var data RotationData
		err := json.Unmarshal([]byte(*message.Data), &data)
		if err != nil {
			log.Fatal(err)
		}
		game.history.Run(&ChangeRotationEvent{message.Time, c.playerId, data.Direction})
	case "fire":
		var data FireData
		err := json.Unmarshal([]byte(*message.Data), &data)
		if err != nil {
			log.Fatal(err)
		}
		game.history.Run(&FireEvent{message.Time, c.playerId, data.ProjectileId, data.Created})
	}

}

func (c *Client) Send(message *Message) {
	c.conn.send <- message
}
