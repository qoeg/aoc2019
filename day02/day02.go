package day02

import (
	"strconv"
)

func compute(intcode []int) []int {
	var opcode int
	pointer := 0

	for {
		opcode = intcode[pointer]
		if opcode == 99 {
			break
		}

		var output int

		switch opcode {
		case 1:
			output = intcode[intcode[pointer+1]] + intcode[intcode[pointer+2]]
		case 2:
			output = intcode[intcode[pointer+1]] * intcode[intcode[pointer+2]]
		default:
			panic("Cannot compute intcode. Unknown operation.")
		}

		intcode[intcode[pointer+3]] = output
		pointer += 4
	}

	return intcode
}

// Answer1 returns the first puzzle answer
func Answer1() string {
	input := make([]int, len(Input))
	copy(input, Input)

	input[1] = 12
	input[2] = 2

	result := compute(input)

	return strconv.Itoa(result[0])
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	var input, result []int

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			input = make([]int, len(Input))
			copy(input, Input)

			input[1] = noun
			input[2] = verb

			result = compute(input)

			if result[0] == 19690720 {
				return strconv.Itoa(100 * noun + verb)
			}
		}
	}

	return "Not found"
}
