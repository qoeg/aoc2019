package day05

import (
	"strconv"
	"github.com/qoeg/aoc2019/computer"
)

// Answer1 returns the first puzzle answer
func Answer1() string {
	c := computer.New(1, computer.Config{})
	c.Input <- 1

	result := c.Run(Input)

	return strconv.Itoa(int(result))
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	c := computer.New(2, computer.Config{})
	c.Input <- 5

	result := c.Run(Input)

	return strconv.Itoa(int(result))
}
