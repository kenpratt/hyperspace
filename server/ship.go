package main

type Ship struct {
	Id       string    `json:"id"`
	Position *Position `json:"position"`
	Angle    float64   `json:"angle"`
}
