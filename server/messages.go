package main

import (
	"encoding/json"
)

// TODO: Organize this into sections for messages, event data, and game object data (probably splitting game objects into separate files, and dropping the Data suffixes).

type Message struct {
	Type string           `json:"type"`
	Time float64          `json:"time"`
	Data *json.RawMessage `json:"data"`
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
	ProjectileId string `json:"projectileId"`
}

type PositionData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
