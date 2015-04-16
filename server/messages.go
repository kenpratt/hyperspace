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

type InitData struct {
	PlayerId  string         `json:"playerId"`
	Constants *GameConstants `json:"constants"`
	State     *UpdateData    `json:"state"`
}

type UpdateData struct {
	Ships       map[string]*Ship       `json:"ships"`
	Projectiles map[string]*Projectile `json:"projectiles"`
}

type FireData struct {
	ProjectileId string `json:"projectileId"`
	Created      uint64 `json:"created"`
}

type AccelerationData struct {
	Direction int8 `json:"direction"`
}

type RotationData struct {
	Direction int8 `json:"direction"`
}
