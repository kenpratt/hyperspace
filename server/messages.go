package main

import (
	"encoding/json"
)

// TODO: Organize this into sections for messages, event data, and game object data (probably splitting game objects into separate files, and dropping the Data suffixes).

type Message struct {
	Type string           `json:"type"`
	Data *json.RawMessage `json:"data"`
}

// Like a message, but for internal use to avoid the extra serialization passes.
// TODO: could we get rid of the event type, and just have FireEvent, MoveEvent, ...? (and switch on type?)
type Event struct {
	Type     string
	PlayerId string
	Data     interface{}
}

type PositionData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Vector struct {
	X float64
	Y float64
}

type InitData struct {
	PlayerId string      `json:"playerId"`
	State    *UpdateData `json:"state"`
}

type UpdateData struct {
	Players     map[string]*PlayerData     `json:"players"`
	Projectiles map[string]*ProjectileData `json:"projectiles"`
}

type PlayerData struct {
	Id       string        `json:"id"`
	Position *PositionData `json:"position"`
	Angle    float64       `json:"angle"`
}

type FireData struct {
	Id   string  `json:"id"`
	Time float64 `json:"time"`
}

type ProjectileData struct {
	Id       string        `json:"id"`
	Position *PositionData `json:"position"`
	Angle    float64       `json:"angle"`
}

func (p *ProjectileData) Vector() *Vector {
	return AngleToVector(p.Angle)
}

func (p *ProjectileData) UpdateOneTick() {
	p.Position = &PositionData{
		X: p.Position.X + p.Vector().X,
		Y: p.Position.Y + p.Vector().Y,
	}
}
