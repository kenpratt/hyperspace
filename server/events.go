package main

type FireEvent struct {
	PlayerId     string
	Time         float64
	ProjectileId string
}

type PositionEvent struct {
	PlayerId string
	Time     float64
	Position *Position
}
