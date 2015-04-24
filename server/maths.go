package main

import (
	"math"
	"math/rand"
	"time"
)

type Point struct {
	X float64 `bson:"x"`
	Y float64 `bson:"y"`
}

type Vector struct {
	X float64 `bson:"x"`
	Y float64 `bson:"y"`
}

func MakePoint(x float64, y float64) *Point {
	return &Point{RoundToPlaces(x, 1), RoundToPlaces(y, 1)}
}

func Round(f float64) float64 {
	return math.Floor(f + 0.5)
}

func RoundToPlaces(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f*shift) / shift
}

func RoundPoint(p *Point) *Point {
	return &Point{RoundToPlaces(p.X, 1), RoundToPlaces(p.Y, 1)}
}

func RoundVector(v *Vector) *Vector {
	return &Vector{RoundToPlaces(v.X, 1), RoundToPlaces(v.Y, 1)}
}

// Converts an angle in degrees between 0 and 359.
func AngleToVector(angle float64) *Vector {
	// Convert to radians.
	r := float64(angle) * 0.01745
	return UnitVector(&Vector{X: math.Sin(r), Y: -math.Cos(r)})
}

func AngleAndSpeedToVector(angle float64, speed uint16) *Vector {
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

func AmountToMove(velocity *Vector, elapsed uint64) (float64, float64) {
	d := float64(elapsed) / 1000
	return velocity.X * d, velocity.Y * d
}

func Random(min int, max int) int {
	d := max - min + 1
	return min + rand.Intn(d)
}

func RandomAngle() float64 {
	return float64(Random(0, 359))
}
