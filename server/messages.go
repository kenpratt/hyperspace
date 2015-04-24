package main

// TODO: Organize this into sections for messages, event data, and game object data (probably splitting game objects into separate files, and dropping the Data suffixes).

//go:generate msgp

type Message struct {
	Type string      `msg:"type"`
	Time uint64      `msg:"time"`
	Data interface{} `msg:"-"`
}

type InitMessage struct {
	Type      string         `msg:"type"`
	Time      uint64         `msg:"time"`
	PlayerId  string         `msg:"playerId"`
	Constants *GameConstants `msg:"constants"`
	State     *GameState     `msg:"state"`
}

type UpdateMessage struct {
	Type        string     `msg:"type"`
	Time        uint64     `msg:"time"`
	State       *GameState `msg:"state"`
	LastEventId uint64     `msg:"lastEvent"`
}

type FireData struct {
	EventId      uint64 `msg:"eventId"`
	ProjectileId string `msg:"projectileId"`
	Created      uint64 `msg:"created"`
}

type AccelerationData struct {
	EventId   uint64 `msg:"eventId"`
	Direction int8   `msg:"direction"`
}

type RotationData struct {
	EventId   uint64 `msg:"eventId"`
	Direction int8   `msg:"direction"`
}
