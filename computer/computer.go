package computer

import (
	"fmt"
)

// Run runs the computer
func Run(intcode []int, input int, verbose bool) (int, []int) {
	pointer := 0
	output := 0

	code := make([]int, len(intcode))
	copy(code, intcode)
	
	for {
		jump := false
		instr := newInstruction(code[pointer], pointer)
		if instr.opcode == 99 {
			break
		}

		switch instr.opcode {
		case 1:
			param1 := instr.parameter(1, code)
			param2 := instr.parameter(2, code)
			code[code[pointer+3]] = param1 + param2
			if verbose {
				fmt.Printf("*%v = %v + %v\n", code[pointer+3], param1, param2)
			}
		case 2:
			param1 := instr.parameter(1, code)
			param2 := instr.parameter(2, code)
			code[code[pointer+3]] = param1 * param2
			if verbose {
				fmt.Printf("*%v = %v * %v\n", code[pointer+3], param1, param2)
			}
		case 3:
			code[code[pointer+1]] = input
			if verbose {
				fmt.Printf("*%v = %v\n", code[pointer+1], input)
			}
		case 4:
			output = instr.parameter(1, code)
			if verbose {
				fmt.Printf("TEST Output = %v (*%v)\n", output, code[pointer+1])
			}
		case 5:
			param1 := instr.parameter(1, code)
			param2 := instr.parameter(2, code)
			if param1 != 0 {
				pointer = param2
				jump = true
			}
		case 6:
			param1 := instr.parameter(1, code)
			param2 := instr.parameter(2, code)
			if param1 == 0 {
				pointer = param2
				jump = true
			}
		case 7:
			param1 := instr.parameter(1, code)
			param2 := instr.parameter(2, code)
			if param1 < param2 {
				code[code[pointer+3]] = 1
			} else {
				code[code[pointer+3]] = 0
			}
		case 8:
			param1 := instr.parameter(1, code)
			param2 := instr.parameter(2, code)
			if param1 == param2 {
				code[code[pointer+3]] = 1
			} else {
				code[code[pointer+3]] = 0
			}
		default:
			panic("Unknown operation.")
		}
		
		if !jump {
			pointer += instr.parameters + 1
		}
	}

	return output, code
}