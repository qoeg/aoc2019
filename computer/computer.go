package computer

import (
	"fmt"
)

type Program struct {
	id int
	pointer int
	code []int
	Input chan int
	Output chan int
	Exit chan int
}

func New(id int, intcode []int) *Program {
	code := make([]int, len(intcode))
	copy(code, intcode)

	return &Program{
		id: id,
		pointer: 0,
		code: code,
		Input: make(chan int, 2),
		Output: make(chan int, 1),
		Exit: make(chan int, 1),
	}
}

// Run runs the computer program
func (c *Program) Run(verbose bool) (result int) {
	for {
		jump := false
		instr := newInstruction(c.code[c.pointer], c.pointer)
		if instr.opcode == 99 {
			c.Exit <- result
			if verbose {
				fmt.Printf("Program %d: Halt opcode hit. Exiting program.\n", c.id)
			}
			break
		}

		switch instr.opcode {
		case 1:
			param1 := instr.parameter(1, c.code)
			param2 := instr.parameter(2, c.code)
			c.code[c.code[c.pointer+3]] = param1 + param2
			if verbose {
				fmt.Printf("Program %d: *%v = %v + %v\n", c.id, c.code[c.pointer+3], param1, param2)
			}
		case 2:
			param1 := instr.parameter(1, c.code)
			param2 := instr.parameter(2, c.code)
			c.code[c.code[c.pointer+3]] = param1 * param2
			if verbose {
				fmt.Printf("Program %d: *%v = %v * %v\n", c.id, c.code[c.pointer+3], param1, param2)
			}
		case 3:
			in := <-c.Input
			c.code[c.code[c.pointer+1]] = in
			if verbose {
				fmt.Printf("Program %d: *%v = %v\n", c.id, c.code[c.pointer+1], in)
			}
		case 4:
			result = instr.parameter(1, c.code)
			if len(c.Output) > 0 {
				<-c.Output
			}
			c.Output <- result
			if verbose {
				fmt.Printf("Program %d: Output = %v (*%v)\n", c.id, result, c.code[c.pointer+1])
			}
		case 5:
			param1 := instr.parameter(1, c.code)
			param2 := instr.parameter(2, c.code)
			if param1 != 0 {
				c.pointer = param2
				jump = true
			}
		case 6:
			param1 := instr.parameter(1, c.code)
			param2 := instr.parameter(2, c.code)
			if param1 == 0 {
				c.pointer = param2
				jump = true
			}
		case 7:
			param1 := instr.parameter(1, c.code)
			param2 := instr.parameter(2, c.code)
			if param1 < param2 {
				c.code[c.code[c.pointer+3]] = 1
			} else {
				c.code[c.code[c.pointer+3]] = 0
			}
		case 8:
			param1 := instr.parameter(1, c.code)
			param2 := instr.parameter(2, c.code)
			if param1 == param2 {
				c.code[c.code[c.pointer+3]] = 1
			} else {
				c.code[c.code[c.pointer+3]] = 0
			}
		default:
			panic("Unknown operation.")
		}
		
		if !jump {
			c.pointer += instr.parameters + 1
		}
	}

	return result
}