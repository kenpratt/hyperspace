// Adapted from https://raw.githubusercontent.com/gorilla/websocket/master/examples/chat/conn.go

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
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
	send    chan *Message
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
		// read raw message off the socket
		_, raw, err := c.ws.ReadMessage()
		if err != nil {
			if err != io.EOF {
				log.Println("Error reading message off socket", err)
			}
			break
		}

		// uncompress LZW
		uncompressed := LzwDecompress(string(raw))

		// unmarshal JSON
		var message Message
		err = json.Unmarshal(uncompressed, &message)
		if err != nil {
			log.Println("Error unmarshalling JSON", err)
			continue
		}

		// process the message
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
func (c *Connection) writeMessage(message *Message) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))

	// marshal to JSON
	str, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshalling message to JSON", err)
		return nil
	}

	// compress with LZW
	compressed := LzwCompress(str)

	// send
	return c.ws.WriteMessage(websocket.TextMessage, []byte(compressed))
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

			if err := c.writeMessage(message); err != nil {
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
	c := &Connection{ws, make(chan *Message, 256), make(chan *Message, 256)}
	makeClient(c)
	go c.writePump()
	c.readPump()
}
