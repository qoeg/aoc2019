package computer

import (
	//"fmt"
)

type instruction struct {
	opcode int
	parameters int
	parameterModes []int
}

func getInstruction(value int) instruction {
	var parameters int
	var parameterModes []int
	
	opcode := value % 100

	switch opcode {
	case 1:
	case 2:
		parameters = 3
		parameterModes = make([]int, 3)
		parameterModes[0] = (value / 100) % 10
		parameterModes[1] = (value / 1000) % 10
		parameterModes[2] = (value / 10000) % 10
	case 3:
	case 4:
		parameters = 1
		parameterModes = make([]int, 1)
		parameterModes[0] = (value / 100) % 10
	case 99:
		parameters = 0
		parameterModes = make([]int, 0)
	default:
		panic("Cannot compute opcode. Unknown operation.")
	}

	return instruction{
		opcode,
		parameters,
		parameterModes,
	}
}

// Run runs the computer
func Run(intcode []int) []int {
	pointer := 0

	for {
		instr := getInstruction(intcode[pointer])
		if instr.opcode == 99 {
			break
		}

		var result int

		switch instr.opcode {
		case 1:
			result = intcode[intcode[pointer+1]] + intcode[intcode[pointer+2]]
			intcode[intcode[pointer+3]] = result
		case 2:
			result = intcode[intcode[pointer+1]] * intcode[intcode[pointer+2]]
			intcode[intcode[pointer+3]] = result
		case 3:
			// Take input and store at intcode[intcode[pointer+1]]
		case 4:
			// Output value stored at intcode[intcode[pointer+1]]
		default:
			panic("Cannot compute opcode. Unknown operation.")
		}
		
		pointer += instr.parameters + 1
	}

	return intcode
}