package day05

import (
	"strconv"
	"github.com/qoeg/aoc2019/computer"
)

// Answer1 returns the first puzzle answer
func Answer1() string {
	c := computer.New(1, Input)
	c.Input <- 1

	result := c.Run(false)

	return strconv.Itoa(result)
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	c := computer.New(2, Input)
	c.Input <- 5

	result := c.Run(false)

	return strconv.Itoa(result)
}
