package day09

import (
	"strconv"
	"github.com/qoeg/aoc2019/computer"
)

// Answer1 returns the first puzzle answer
func Answer1() string {
	p := computer.New(1, computer.Config{
		MinMemoryAlloc: 100000,
		Print: false,
	})
	p.Input <- 1

	result := p.Run64(Input)

	return strconv.FormatInt(result, 10)
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	p := computer.New(2, computer.Config{
		MinMemoryAlloc: 100000,
		Print: false,
	})
	p.Input <- 2

	result := p.Run64(Input)

	return strconv.FormatInt(result, 10)
}
