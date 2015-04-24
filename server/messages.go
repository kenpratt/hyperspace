package main

import (
	"encoding/json"
)

// TODO: Organize this into sections for messages, event data, and game object data (probably splitting game objects into separate files, and dropping the Data suffixes).

type Message struct {
	Type string           `bson:"type"`
	Time uint64           `bson:"time"`
	Data *json.RawMessage `bson:"data"`
}

type InitMessage struct {
	Type      string         `bson:"type"`
	Time      uint64         `bson:"time"`
	PlayerId  string         `bson:"playerId"`
	Constants *GameConstants `bson:"constants"`
	State     *GameState     `bson:"state"`
}

type UpdateMessage struct {
	Type        string     `bson:"type"`
	Time        uint64     `bson:"time"`
	State       *GameState `bson:"state"`
	LastEventId uint64     `bson:"lastEvent"`
}

type FireData struct {
	EventId      uint64 `bson:"eventId"`
	ProjectileId string `bson:"projectileId"`
	Created      uint64 `bson:"created"`
}

type AccelerationData struct {
	EventId   uint64 `bson:"eventId"`
	Direction int8   `bson:"direction"`
}

type RotationData struct {
	EventId   uint64 `bson:"eventId"`
	Direction int8   `bson:"direction"`
}
