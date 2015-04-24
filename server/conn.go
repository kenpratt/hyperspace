// Adapted from https://raw.githubusercontent.com/gorilla/websocket/master/examples/chat/conn.go
//
// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channels of inbound and outbound messages.
	receive chan *Message
	send    chan interface{}
}

// readPump pumps messages from the websocket connection to the hub.
func (c *Connection) readPump() {
	defer func() {
		c.ws.Close()
		close(c.receive)
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var message Message
		err := c.ws.ReadJSON(&message)
		if err != nil {
			break
		}
		if message.Type == "h" {
			// respond to heartbeat right away, but still send it to the client as well
			c.send <- &Message{Type: "h", Time: MakeTimestamp()}
		}
		c.receive <- &message
	}
}

// write a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// write a JSON message.
func (c *Connection) writeJSON(message interface{}) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	data, err := bson.Marshal(message)
	if err != nil {
		panic(err)
	}
	return c.ws.WriteMessage(websocket.BinaryMessage, data)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.writeJSON(message); err != nil {
				log.Println("Error sending message:", err)
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				log.Println("Error sending ping:", err)
				return
			}
		}
	}
}

// serverWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := &Connection{ws, make(chan *Message, 256), make(chan interface{}, 256)}
	makeClient(c)
	go c.writePump()
	c.readPump()
}
