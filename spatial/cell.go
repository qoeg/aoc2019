package spatial

import (
	"math"
)

type cell struct {
	index int
	mark rune
	pos Coordinate
}

// Cell represents a 2D location in a grid
type Cell struct {
	cell
}

// NewCell creates a new Cell instance with a rune and position
func NewCell(mark rune, pos Coordinate) Cell {
	return Cell{cell{0, mark, pos}}
}

// NewEmptyCell creates a new Cell with only a position
func NewEmptyCell(pos Coordinate) Cell {
	return Cell{cell{0, rune(' '), pos}}
}

// Mark represents the cell in text formats
func (c Cell) Mark() rune {
	return c.mark
}

// Pos is the cell's coordinate position
func (c Cell) Pos() Coordinate {
	return c.pos
}

// String is the cell's mark as a sting
func (c Cell) String() string {
	return string(c.mark)
}

type Path []Cell

// MinX finds the minimum X value from a slice of Coordinates
func (p Path) MinX() int {
	min := math.MaxInt32
	for _, c := range []Cell(p) {
		if c.pos.X < min {
			min = c.pos.X
		}
	}
	return min
}

// MaxX finds the maximum X value from a slice of Coordinates
func (p Path) MaxX() int {
	max := 0
	for _, c := range []Cell(p) {
		if c.pos.X > max {
			max = c.pos.X
		}
	}
	return max
}

// MinY finds the minimum Y value from a slice of Coordinates
func (p Path) MinY() int {
	min := math.MaxInt32
	for _, c := range []Cell(p) {
		if c.pos.Y < min {
			min = c.pos.Y
		}
	}
	return min
}

// MaxY finds the maximum Y value from a slice of Coordinates
func (p Path) MaxY() int {
	max := 0
	for _, c := range []Cell(p) {
		if c.pos.Y > max {
			max = c.pos.Y
		}
	}
	return max
}
