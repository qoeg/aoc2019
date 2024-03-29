package spatial

import (
	"fmt"

	"github.com/qoeg/aoc2019/util"
)

// Grid represents a two dimensional grid of cells
type Grid [][]Cell

// Draw updates cells in the Grid with new cells
func (g Grid) Draw(line []Cell) {
	for i := range line {
		coord := line[i].Pos()
		if i > 0 && g[coord.X][coord.Y].Mark() != '.' {
			if found, _ := util.Contains(line, line[i]); !found {
				line[i].mark = 'X'
			}
		}

		g[coord.X][coord.Y] = line[i]
	}
}

// NewGrid initializes a new Grid with the specified dimensions
func NewGrid(width, height int) (g Grid) {
	g = make([][]Cell, width)
	for x := 0; x < width; x++ {
		g[x] = make([]Cell, height)
		for y := 0; y < height; y++ {
			g[x][y] = NewCell('.', Coordinate{x, y})
		}
	}

	return g
}

// Clone creates a new clone of the Grid
func (g Grid) Clone() Grid {
	clone := [][]Cell{}
	for x := 0; x < len(g); x++ {
		clone = append(clone, []Cell{})
		for y := 0; y < len(g[x]); y++ {
			clone[x] = append(clone[x], g[x][y])
		}
	}
	return clone
}

// Neighbors finds the four adjacent cells for the Coordinate.
// The optional include parameter is a rune whitelist for if only certain runes are wanted.
func (g Grid) Neighbors(c Coordinate, include ...rune) []Cell {
	n := []Cell{}
	if len(g) == 0 || len(g[0]) == 0 {
		fmt.Println("grid.Neighbors(): the grid is empty")
		return n
	}

	pos := c.Pos()
	sides := make([]Cell, 4)

	if pos.Y != 0 {
		sides[0] = g[pos.X][pos.Y-1]
	}
	if pos.X != 0 {
		sides[1] = g[pos.X-1][pos.Y]
	}
	if pos.X != len(g)-1 {
		sides[2] = g[pos.X+1][pos.Y]
	}
	if pos.Y != len(g[pos.X])-1 {
		sides[3] = g[pos.X][pos.Y+1]
	}

	for _, s := range sides {
		if len(include) == 0 {
			n = append(n, s)
			continue
		}
		for _, in := range include {
			if s.mark == in {
				n = append(n, s)
				break
			}
		}
	}

	return n
}

// Print will print the grid to the console, and 
// replace any cells with the provided cell masks
func (g Grid) Print(masks ...Cell) {
	for y := 0; y < len(g[0]); y++ {
		for x := 0; x < len(g); x++ {
			// print any objects for the location
			masked := false
			for i, o := range masks {
				if (o.Pos() == Coordinate{x, y}) {
					fmt.Print(o)
					masks = append(masks[:i], masks[i+1:]...)
					masked = true
					break
				}
			}

			if !masked {
				fmt.Print(string(g[x][y].mark))
			}
		}
		fmt.Println()
	}
}

// With creates a clone of the Grid with the provided cells as replacements
func (g Grid) With(cells []Cell) Grid {
	clone := g.Clone()
	for _, c := range cells {
		clone[c.Pos().X][c.Pos().Y] = c
	}
	return clone
}

// Parse creates a two dimensional slice of Cells from a grid string 
func Parse(input string) (grid [][]Cell) {
	return ParseOut(input, nil)
}

// ParseOut creates a two dimensional slice of Cells from a grid string.
// The replace function parameter can be used to replace certain cells with a new mark.
// The replace parameter can be nil.
func ParseOut(input string, replace func(Cell) rune) (grid [][]Cell) {
	var x, y int
	for _, r := range input {
		if r == rune(10) {
			y++
			x = 0
			continue
		}

		c := NewCell(r, Coordinate{X:x, Y:y})
		if replace != nil {
			c.mark = replace(c)
		}

		if y == 0 {
			grid = append(grid, []Cell{})
		}

		grid[x] = append(grid[x], c)
		x++
	}

	return grid
}
