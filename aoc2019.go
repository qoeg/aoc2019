package main

import (
	"fmt"

	"github.com/qoeg/aoc2019/day01"
	"github.com/qoeg/aoc2019/day02"
	"github.com/qoeg/aoc2019/day03"
	"github.com/qoeg/aoc2019/day04"
)

type day struct {
	key int
	ans1 func() string
	ans2 func() string
}

var days = []day{
	{1, day01.Answer1, day01.Answer2},
	{2, day02.Answer1, day02.Answer2},
	{3, day03.Answer1, day03.Answer2},
	{4, day04.Answer1, day04.Answer2},
}

func main() {
	//TODO: add CLI for selecting days
	selected := []int{}

	var d day
	var found bool
	for _, d = range days {
		for _, i := range selected {
			if d.key == i {
				run(d)
				found = true
			}
		}
	}

	// none selected, run latest
	if !found {
		run(d)
	}
}

func run(d day) {
	fmt.Printf("------- Day %d -------\n", d.key)
	fmt.Printf("Puzzle 1 Answer: %s\n", d.ans1())
	fmt.Printf("Puzzle 2 Answer: %s\n", d.ans2())
}
