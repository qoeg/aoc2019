package day03

import (
	"fmt"
	"math"
	"strconv"

	"github.com/qoeg/aoc2019/grid"
)

var intersections []*intersection

type intersection struct {
	coord grid.Coordinate
	l1steps int
	l2steps int
}

func getLine(start grid.Coordinate, ops []string) []grid.Cell {
	cursor := start
	line := []grid.Cell{grid.NewCell('o', grid.Coordinate{X:cursor.X, Y:cursor.Y})}
	index := 1

	for k, op := range ops {
		dir := op[0]

		dist, err := strconv.Atoi(op[1:])
		if err != nil {
			fmt.Println("Could not parse distance integer", err)
			break
		}

		line = append(line, make([]grid.Cell, dist)...)

		if k > 0 {
			line[index-1] = grid.NewCell('+', grid.Coordinate{X:cursor.X, Y:cursor.Y})
		}

		for count := 0; count < dist; count++ {
			var mark rune
			switch dir {
			case 'L':
				cursor.X--
				mark = '-'
			case 'U':
				cursor.Y--
				mark = '|'
			case 'R':
				cursor.X++
				mark = '-'
			case 'D':
				cursor.Y++
				mark = '|'
			}

			line[index] = grid.NewCell(mark, grid.Coordinate{X:cursor.X, Y:cursor.Y})
			index++
		}
	}

	return line
}

func getIntersections(start grid.Coordinate, line1, line2 []grid.Cell) []*intersection {
	lookup := map[grid.Coordinate]*intersection{}
	intrs := []*intersection{}

	for i := range line1 {
		pos := line1[i].Pos()
		if pos == start {
			continue
		}

		lookup[pos] = &intersection{coord: pos, l1steps: i}
	}

	for j := range line2 {
		pos := line2[j].Pos()
		
		if lookup[pos] != nil {
			lookup[pos].l2steps = j
			intrs = append(intrs, lookup[pos])
		}
	}

	return intrs
}

func getMinDistance(start grid.Coordinate, intrs []*intersection) int {
	min := math.MaxInt32
	for _, intr := range intrs {
		dist := grid.Distance(start, intr.coord)
		if dist < min {
			min = dist
		}
	}
	return min
}

func test1() int {
	start := grid.Coordinate{X:0, Y:0}
	line1 := getLine(start, []string{"R8","U5","L5","D3"})
	line2 := getLine(start, []string{"U7","R6","D4","L4"})
	intersections = getIntersections(start, line1, line2)
	
	return getMinDistance(start, intersections)
}

func test2() int {
	start := grid.Coordinate{X:0, Y:0}
	line1 := getLine(start, []string{"R75","D30","R83","U83","L12","D49","R71","U7","L72"})
	line2 := getLine(start, []string{"U62","R66","U55","R34","D71","R55","D58","R83"})
	intersections = getIntersections(start, line1, line2)

	return getMinDistance(start, intersections)
}

func test3() int {
	start := grid.Coordinate{X:0, Y:0}
	line1 := getLine(start, []string{"R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"})
	line2 := getLine(start, []string{"U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"})
	intersections = getIntersections(start, line1, line2)

	return getMinDistance(start, intersections)
}

// Answer1 returns the first puzzle answer
func Answer1() string {
	start := grid.Coordinate{X:0, Y:0}
	line1 := getLine(start, Wire1)
	line2 := getLine(start, Wire2)
	intersections = getIntersections(start, line1, line2)

	return strconv.Itoa(getMinDistance(start, intersections))
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	min := math.MaxInt32
	for _, intr := range intersections {
		sum := intr.l1steps + intr.l2steps
		if sum < min {
			min = sum
		}
	}

	return strconv.Itoa(min)
}
