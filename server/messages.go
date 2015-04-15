package main

import (
	"encoding/json"
)

type Message struct {
	Type string           `json:"type"`
	Data *json.RawMessage `json:"data"`
}

type InitData struct {
	Id uint16 `json:"id"`
}

type PositionData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PlayerData struct {
	Id uint16  `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
}

type FireData struct {
	Id   string  `json:"id"`
	Time float64 `json:"time"`
}

type ProjectileData struct {
	Id       string       `json:"id"`
	Position PositionData `json:"position"`
	Angle    float64      `json:"angle"`
}
