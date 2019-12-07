package spatial

import (
	"math"
)

// Coordinate represents a position in a x,y coordinate system
type Coordinate struct {
	X int
	Y int
}

// Pos returns the Coordinate position
func (c Coordinate) Pos() Coordinate {
	return c
}

// Distance gets the Manhatten distance between two coordinates
func Distance(coord1, coord2 Coordinate) int {
	return int(math.Abs(float64(coord1.X - coord2.X)) + math.Abs(float64(coord1.Y - coord2.Y)))
}

// MinX finds the minimum X value from a slice of Coordinates
func MinX(coords []Coordinate) int {
	min := math.MaxInt32
	for _, c := range coords {
		if c.X < min {
			min = c.X
		}
	}
	return min
}

// MaxX finds the maximum X value from a slice of Coordinates
func MaxX(coords []Coordinate) int {
	max := 0
	for _, c := range coords {
		if c.X > max {
			max = c.X
		}
	}
	return max
}

// MinY finds the minimum Y value from a slice of Coordinates
func MinY(coords []Coordinate) int {
	min := math.MaxInt32
	for _, c := range coords {
		if c.Y < min {
			min = c.Y
		}
	}
	return min
}

// MaxY finds the maximum Y value from a slice of Coordinates
func MaxY(coords []Coordinate) int {
	max := 0
	for _, c := range coords {
		if c.Y > max {
			max = c.Y
		}
	}
	return max
}
