package main

import (
	"fmt"

	"github.com/qoeg/aoc2019/day01"
	"github.com/qoeg/aoc2019/day02"
)

type day struct {
	key int
	fun func()
}

var days = []day{
	{1, d1},
	{2, d2},
}

func main() {
	//TODO: add CLI for selecting days
	selected := []int{}

	var d day
	var found bool
	for _, d = range days {
		for _, i := range selected {
			if d.key == i {
				fmt.Printf("Day %d\n", d.key)
				d.fun()
				found = true
			}
		}
	}

	if !found {
		fmt.Printf("Day %d\n", d.key)
		d.fun()
	}
}

func d2() {
	fmt.Printf("Puzzle 1 Answer: %s\n", day02.Answer1())
	fmt.Printf("Puzzle 2 Answer: %s\n", day02.Answer2())
}

func d1() {
	fmt.Printf("Puzzle 1 Answer: %s\n", day01.Answer1())
	fmt.Printf("Puzzle 2 Answer: %s\n", day01.Answer2())
}
