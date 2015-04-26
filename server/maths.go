package main

import (
	"math"
	"math/rand"
	"time"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
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
	r := angle * 0.01745
	return UnitVector(&Vector{X: math.Sin(r), Y: -math.Cos(r)})
}

func AngleAndSpeedToVector(angle float64, speed float64) *Vector {
	return MultiplyVector(AngleToVector(angle), speed)
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

func MultiplyVector(vector *Vector, f float64) *Vector {
	return &Vector{
		X: vector.X * f,
		Y: vector.Y * f,
	}
}

func AddVectors(vector1 *Vector, vector2 *Vector) *Vector {
	return &Vector{
		X: vector1.X + vector2.X,
		Y: vector1.Y + vector2.Y,
	}
}

func AddVectorToPoint(vector *Vector, point *Point) *Point {
	return &Point{
		X: point.X + vector.X,
		Y: point.Y + vector.Y,
	}
}

func MakeTimestamp() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}

func Random(min int, max int) int {
	d := max - min + 1
	return min + rand.Intn(d)
}

func RandomFloat(min float64, max float64) float64 {
	d := max - min + 1
	return min + rand.Float64()*d
}

func RandomAngle() float64 {
	return float64(Random(0, 359))
}

func DistanceBetweenPoints(p1 *Point, p2 *Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func CalculateCenter(points []*Point) *Point {
	var sx float64 = 0
	var sy float64 = 0
	for _, p := range points {
		sx += p.X
		sy += p.Y
	}
	return &Point{
		X: sx / float64(len(points)),
		Y: sy / float64(len(points)),
	}
}

func IsColliding(p1 *Point, r1 float64, p2 *Point, r2 float64) bool {
	return DistanceBetweenPoints(p1, p2) < (r1 + r2)
}
