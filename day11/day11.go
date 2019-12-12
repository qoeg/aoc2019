package day11

import (
	"strconv"
	"github.com/qoeg/aoc2019/computer"
	"github.com/qoeg/aoc2019/spatial"
)

// Answer1 returns the first puzzle answer
func Answer1() string {
	painted := make(map[spatial.Coordinate]int)
	program := computer.New(1, computer.Config{
		MinMemoryAlloc: 5000,
		Print: false,
	})
	robot := newRobot(program, spatial.Coordinate{X:50, Y:50})
	robot.run(painted)

	return strconv.Itoa(len(painted))
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	painted := make(map[spatial.Coordinate]int)
	program := computer.New(2, computer.Config{
		MinMemoryAlloc: 5000,
		Print: false,
	})
	robot := newRobot(program, spatial.Coordinate{X:0, Y:0})

	painted[robot.position] = 1
	robot.run(painted)

	grid := spatial.NewGrid(41, 6)
	cells := make([]spatial.Cell, len(painted))

	var i int
	for coord, color := range painted {
		if color == 1 {
			cells[i] = spatial.NewCell('#', coord)
		} else {
			cells[i] = spatial.NewCell('.', coord)
		}
		i++
	}

	grid.Print(cells...)

	return "N/A (image)"
}
