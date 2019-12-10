package computer

import (
	"fmt"
)

// Program represents an intcode computer program
type Program struct {
	id int
	pointer int64
	relBase int64
	memory []int64
	config Config
	Input chan int64
	Output chan int64
	Exit chan int64
}

// Config allows setting certain configurations on a program
type Config struct {
	MinMemoryAlloc int
	Print bool
}

// New creates a new intcode computer program with an ID
// and configurations for running the program
func New(id int, config Config) *Program {
	return &Program{
		id: id,
		pointer: 0,
		relBase: 0,
		memory: []int64{},
		config: config,
		Input: make(chan int64, 2),
		Output: make(chan int64, 1),
		Exit: make(chan int64, 1),
	}
}

// Run runs the computer program with 32-bit intcode
func (p *Program) Run(intcode []int) (result int) {
	memorySize := len(intcode)
	if p.config.MinMemoryAlloc > memorySize {
		memorySize = p.config.MinMemoryAlloc
	}

	p.memory = make([]int64, memorySize)
	for i, code := range intcode {
		p.memory[i] = int64(code)
	}

	return int(run(p))
}

// Run64 runs the computer program with 64-bit intcode
func (p *Program) Run64(intcode []int64) (result int64) {
	memorySize := len(intcode)
	if p.config.MinMemoryAlloc > memorySize {
		memorySize = p.config.MinMemoryAlloc
	}

	p.memory = make([]int64, memorySize)
	copy(p.memory, intcode)

	return run(p)
}

func run(p *Program) (result int64) {
	for {
		jump := false
		instr := newInstruction(p, p.memory[p.pointer])
		if instr.opcode == 99 {
			p.Exit <- result
			if p.config.Print {
				fmt.Printf("Program %d: Halt: Exiting program.\n", p.id)
			}
			break
		}

		switch instr.opcode {
		case 1:
			param1 := instr.read(1)
			param2 := instr.read(2)
			pos := instr.write(3, param1 + param2)
			if p.config.Print {
				fmt.Printf("Program %d: Add: *%v = %v + %v\n", p.id, pos, param1, param2)
			}
		case 2:
			param1 := instr.read(1)
			param2 := instr.read(2)
			pos := instr.write(3, param1 * param2)
			if p.config.Print {
				fmt.Printf("Program %d: Mult: *%v = %v * %v\n", p.id, pos, param1, param2)
			}
		case 3:
			in := <-p.Input
			pos := instr.write(1, in)
			if p.config.Print {
				fmt.Printf("Program %d: Input: *%v = %v\n", p.id, pos, in)
			}
		case 4:
			result = instr.read(1)
			if len(p.Output) > 0 {
				<-p.Output
			}
			p.Output <- result
			if p.config.Print {
				fmt.Printf("Program %d: Output: %v\n", p.id, result)
			}
		case 5:
			param1 := instr.read(1)
			param2 := instr.read(2)
			if param1 != 0 {
				p.pointer = param2
				jump = true
				if p.config.Print {
					fmt.Printf("Program %d: Jump: %v\n", p.id, param2)
				}
			}
		case 6:
			param1 := instr.read(1)
			param2 := instr.read(2)
			if param1 == 0 {
				p.pointer = param2
				jump = true
				if p.config.Print {
					fmt.Printf("Program %d: Jump: %v\n", p.id, param2)
				}
			}
		case 7:
			param1 := instr.read(1)
			param2 := instr.read(2)
			if param1 < param2 {
				pos := instr.write(3, 1)
				if p.config.Print {
					fmt.Printf("Program %d: LT: *%v = (%v < %v) // =%v\n", p.id, pos, param1, param2, 1)
				}
			} else {
				pos := instr.write(3, 0)
				if p.config.Print {
					fmt.Printf("Program %d: LT: *%v = (%v < %v) // =%v\n", p.id, pos, param1, param2, 0)
				}
			}
		case 8:
			param1 := instr.read(1)
			param2 := instr.read(2)
			if param1 == param2 {
				pos := instr.write(3, 1)
				if p.config.Print {
					fmt.Printf("Program %d: EQ: *%v = (%v == %v) // =%v\n", p.id, pos, param1, param2, 1)
				}
			} else {
				pos := instr.write(3, 0)
				if p.config.Print {
					fmt.Printf("Program %d: EQ: *%v = (%v == %v) // =%v\n", p.id, pos, param1, param2, 0)
				}
			}
		case 9:
			param1 := instr.read(1)
			p.relBase += param1
			if p.config.Print {
				fmt.Printf("Program %d: RB: += %v // =%v\n", p.id, param1, p.relBase)
			}
		default:
			panic("Unknown operation.")
		}
		
		if !jump {
			p.pointer += int64(instr.parameters + 1)
		}
	}

	return result
}
