package main

import (
	"math"
	"math/rand"
	"time"
)

type Coordinate struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Angle uint16

// Converts an angle in degrees between 0 and 360.
func AngleToVector(angle Angle) *Vector {
	// Convert to radians.
	r := float64(angle) * 0.01745
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

func MakeTimestamp() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}

func AmountToRotate(direction int8, speed uint16, elapsed uint64) Angle {
	d := float64(speed) * float64(elapsed) / 1000
	return Angle(float64(direction) * d)
}

func AmountToMove(angle Angle, speed uint16, elapsed uint64) (int64, int64) {
	v := AngleToVector(angle)
	d := float64(speed) * float64(elapsed) / 1000
	return int64(v.X * d), int64(v.Y * d)
}

func Random(min int, max int) int {
	d := max - min + 1
	return min + rand.Intn(d)
}

func RandomAngle() Angle {
	return Angle(Random(0, 359))
}
