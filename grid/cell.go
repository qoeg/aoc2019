package grid

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
