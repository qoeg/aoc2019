package main

import (
	"fmt"
)

type day struct {
	key int
	ans1 func() string
	ans2 func() string
}

func (d day) run() {
	fmt.Printf("------- Day %d -------\n", d.key)
	fmt.Printf("Puzzle 1 Answer: %s\n", d.ans1())
	fmt.Printf("Puzzle 2 Answer: %s\n", d.ans2())
}
