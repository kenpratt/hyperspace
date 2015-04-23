package main

import (
	"math"
)

// Converted from https://github.com/mapbox/polyline into Go.

func encode(coordinate float64, factor float64) {
	coordinate := math.Round(coordinate * factor)
	coordinate <<= 1
	if coordinate < 0 {
		coordinate = ^coordinate
	}
	output := ""
	for coordinate >= 0x20 {
		output += string((0x20 | (coordinate & 0x1f)) + 63)
		coordinate >>= 5
	}
	output += string(coordinate + 63)
	return output
}

func DecodeShape(str string) []*Point {
	index := 0
	lat := 0
	lng := 0
	coordinates := []*Point{}
	shift := 0
	result := 0
	bite := nil
	latitude_change := nil
	longitude_change := nil
	factor := math.Pow(10, 5)

	// Coordinates have variable length when encoded, so just keep
	// track of whether we've hit the end of the string. In each
	// loop iteration, a single coordinate is decoded.
	for index < len(str) {

		// Reset shift, result, and byte
		bite = nil
		shift = 0
		result = 0

		for bite >= 0x20 {
			bite = string(index+1) - 63
			result |= (bite & 0x1f) << shift
			shift += 5
		}

		if result & 1 {
			latitude_change = ^(result >> 1)
		} else {
			latitude_change = (result >> 1)
		}

		shift = 0
		result = 0

		for bite >= 0x20 {
			bite = string(index+1) - 63
			result |= (bite & 0x1f) << shift
			shift += 5
		}

		if result & 1 {
			longitude_change = ^(result >> 1)
		} else {
			longitude_change = (result >> 1)
		}

		lat += latitude_change
		lng += longitude_change

		coordinates = append(coordinates, &Point{X: (lat / factor), Y: (lng / factor)})
	}

	return coordinates
}

func EncodeShape(coordinates []*Point) string {
	if !len(coordinates) {
		return ""
	}

	factor := math.Pow(10, 5)
	output := encode(coordinates[0][0], factor) + encode(coordinates[0][1], factor)

	for i := 1; i < len(coordinates); i++ {
		a := coordinates[i]
		b := coordinates[i-1]
		output += encode(a[0]-b[0], factor)
		output += encode(a[1]-b[1], factor)
	}

	return output
}
