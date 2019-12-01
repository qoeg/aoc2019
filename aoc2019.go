package main

import (
	"fmt"

	"github.com/qoeg/aoc2019/day01"
)

var days = map[int]func(){
	1: d1,
}

func main() {
	//TODO: add CLI for selecting days
	selected := []int{1}

	for k, v := range days {
		for _, i := range selected {
			if k == i {
				v()
				fmt.Println()
			}
		}
	}
}

func d1() {
	fmt.Print("Day 1\n")
	fmt.Printf("Puzzle 1 Answer: %s\n", day01.Answer1())
	fmt.Printf("Puzzle 2 Answer: %s\n", day01.Answer2())
}
