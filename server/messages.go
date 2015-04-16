package main

import (
	"encoding/json"
)

// TODO organize this into sections for messages, event data, and game object data (probably splitting game objects into separate files, and dropping the Data suffixes)

type Message struct {
	Type string           `json:"type"`
	Data *json.RawMessage `json:"data"`
}

// like a message, but for internal use to avoid the extra serialization passes
// TODO could we get rid of the event type, and just have FireEvent, MoveEvent, ...? (and switch on type?)
type Event struct {
	Type     string
	PlayerId string
	Data     interface{}
}

type InitData struct {
	PlayerId string      `json:"playerId"`
	State    *UpdateData `json:"state"`
}

type UpdateData struct {
	Ships       map[string]*Ship       `json:"ships"`
	Projectiles map[string]*Projectile `json:"projectiles"`
}

type FireData struct {
	Id   string  `json:"id"`
	Time float64 `json:"time"`
}

type PositionData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
