package main

import (
	"encoding/json"
)

// TODO: Organize this into sections for messages, event data, and game object data (probably splitting game objects into separate files, and dropping the Data suffixes).

type Message struct {
	Type string           `json:"type"`
	Time uint64           `json:"time"`
	Data *json.RawMessage `json:"data"`
}

type InitMessage struct {
	Type      string         `json:"type"`
	Time      uint64         `json:"time"`
	PlayerId  string         `json:"playerId"`
	Constants *GameConstants `json:"constants"`
	State     *GameState     `json:"state"`
}

type UpdateMessage struct {
	Type        string     `json:"type"`
	Time        uint64     `json:"time"`
	State       *GameState `json:"state"`
	LastEventId uint64     `json:"lastEvent"`
}

type FireData struct {
	EventId      uint64 `json:"eventId"`
	ProjectileId string `json:"projectileId"`
	Created      uint64 `json:"created"`
}

type AccelerationData struct {
	EventId   uint64 `json:"eventId"`
	Direction int8   `json:"direction"`
}

type RotationData struct {
	EventId   uint64 `json:"eventId"`
	Direction int8   `json:"direction"`
}
