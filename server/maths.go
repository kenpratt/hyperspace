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

func AddFloatToAngle(a Angle, f float64) Angle {
	f += float64(a)
	for f < 0 {
		f += 360
	}
	if f >= 360 {
		return Angle(int64(f) % 360)
	} else {
		return Angle(f)
	}
}

// Converts an angle in degrees between 0 and 359.
func AngleToVector(angle Angle) *Vector {
	// Convert to radians.
	r := float64(angle) * 0.01745
	return UnitVector(&Vector{X: math.Sin(r), Y: -math.Cos(r)})
}

func AngleAndSpeedToVector(angle Angle, speed uint16) *Vector {
	return MultiplyVector(AngleToVector(angle), int(speed))
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

func MultiplyVector(vector *Vector, a int) *Vector {
	return &Vector{
		X: (vector.X * float64(a)),
		Y: (vector.Y * float64(a)),
	}
}

func MakeTimestamp() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}

func AmountToRotate(direction int8, speed uint16, elapsed uint64) float64 {
	d := float64(speed) * float64(elapsed) / 1000
	return float64(direction) * d
}

func AmountToMove(velocity *Vector, elapsed uint64) (int64, int64) {
	d := float64(elapsed) / 1000
	return int64(velocity.X * d), int64(velocity.Y * d)
}

func Random(min int, max int) int {
	d := max - min + 1
	return min + rand.Intn(d)
}

func RandomAngle() Angle {
	return Angle(Random(0, 359))
}
