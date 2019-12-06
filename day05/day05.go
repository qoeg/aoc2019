package day05

import (
	"strconv"
	"github.com/qoeg/aoc2019/computer"
)

// Answer1 returns the first puzzle answer
func Answer1() string {
	output, _ := computer.Run(Input, 1, false)
	return strconv.Itoa(output)
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	output, _ := computer.Run(Input, 5, false)
	return strconv.Itoa(output)
}
