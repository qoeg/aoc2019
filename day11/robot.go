package day11

import (
	"github.com/qoeg/aoc2019/computer"
	"github.com/qoeg/aoc2019/spatial"
)

type robot struct {
	position spatial.Coordinate
	heading direction
	program *computer.Program
}

func newRobot(program *computer.Program, position spatial.Coordinate) *robot {
	return &robot{
		position: position,
		heading: up,
		program: program,
	}
}

func (r *robot) move(input int) {
	r.heading = r.heading.turn(input)

	switch r.heading {
	case up:
		r.position.Y--
	case right:
		r.position.X++
	case down:
		r.position.Y++
	case left:
		r.position.X--
	}
}

func (r *robot) run(painted map[spatial.Coordinate]int) {
	go r.program.Run64(Input)

	func() {
		for {
			if color, ok := painted[r.position]; ok {
				r.program.Input <- int64(color)
			} else {
				r.program.Input <- 0
			}

			select {
			case paint := <-r.program.Output:
				painted[r.position] = int(paint)
				change := <-r.program.Output
				r.move(int(change))
			case <-r.program.Exit:
				return
			}
		}
	}()
}
