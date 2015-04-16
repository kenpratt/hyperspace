package main

import (
	"math"
	"time"
)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Converts an angle in degrees between 0 and 360.
func AngleToVector(angle float64) *Vector {
	// Convert to radians.
	r := angle * 0.01745
	return UnitVector(&Vector{X: math.Sin(r), Y: -math.Cos(r)})
}

func Magnitude(vector *Vector) float64 {
	return math.Sqrt(vector.X*vector.X + vector.Y*vector.Y)
}

func UnitVector(vector *Vector) *Vector {
	return &Vector{
		X: (vector.X / Magnitude(vector)),
		Y: (vector.Y / Magnitude(vector)),
	}
}

func makeTimestamp() float64 {
	return float64(time.Now().UnixNano() / int64(time.Millisecond))
}
