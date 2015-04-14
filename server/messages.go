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
