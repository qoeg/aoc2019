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
		Output: make(chan int64, 100),
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

	return int(p.exec())
}

// Run64 runs the computer program with 64-bit intcode
func (p *Program) Run64(intcode []int64) (result int64) {
	memorySize := len(intcode)
	if p.config.MinMemoryAlloc > memorySize {
		memorySize = p.config.MinMemoryAlloc
	}

	p.memory = make([]int64, memorySize)
	copy(p.memory, intcode)

	return p.exec()
}

func (p *Program) exec() (result int64) {
	for {
		instr := newInstruction(p.memory[p.pointer])

		if p.config.Print {
			fmt.Printf("Program %d: @ %v, opcode = %v\n", p.id, p.pointer, instr.opcode)
		}

		if instr.opcode == 99 {
			break;
		}

		switch instr.opcode {
		case 1:
			p.add(instr)
		case 2:
			p.mult(instr)
		case 3:
			value := <-p.Input
			p.input(instr, value)
		case 4:
			result = p.output(instr)
			p.Output <- result
		case 5:
			p.jump(instr, func(v int64) bool {return v != 0})
		case 6:
			p.jump(instr, func(v int64) bool {return v == 0})
		case 7:
			p.comp(instr, func(v1, v2 int64) bool {return v1 < v2})
		case 8:
			p.comp(instr, func(v1, v2 int64) bool {return v1 == v2})
		case 9:
			p.rebase(instr)
		default:
			panic("Unknown operation.")
		}
	}

	p.halt(result)
	return result
}

func (p *Program) add(in instruction) {
	val1 := in.read(p, 1)
	val2 := in.read(p, 2)
	addr := in.write(p, 3, val1 + val2)
	p.pointer += in.len()

	if p.config.Print {
		fmt.Printf("Program %d: Add: *%v = %v + %v\n", p.id, addr, val1, val2)
	}
}

func (p *Program) mult(in instruction) {
	val1 := in.read(p, 1)
	val2 := in.read(p, 2)
	addr := in.write(p, 3, val1 * val2)
	p.pointer += in.len()

	if p.config.Print {
		fmt.Printf("Program %d: Mult: *%v = %v * %v\n", p.id, addr, val1, val2)
	}
}

func (p *Program) input(in instruction, value int64) {
	pos := in.write(p, 1, value)
	p.pointer += in.len()

	if p.config.Print {
		fmt.Printf("Program %d: Input: *%v = %v\n", p.id, pos, value)
	}
}

func (p *Program) output(in instruction) (result int64) {
	result = in.read(p, 1)
	p.pointer += in.len()

	if p.config.Print {
		fmt.Printf("Program %d: Output: %v\n", p.id, result)
	}

	return result
}

func (p *Program) jump(in instruction, predicate func(int64) bool) {
	val1 := in.read(p, 1)
	val2 := in.read(p, 2)

	if predicate(val1) {
		p.pointer = val2
		if p.config.Print {
			fmt.Printf("Program %d: JUMP: %v\n", p.id, val2)
		}
	} else {
		p.pointer += in.len()
	}
}

func (p *Program) comp(in instruction, comparator func(int64, int64) bool) {
	val1 := in.read(p, 1)
	val2 := in.read(p, 2)

	var res = 0
	if comparator(val1, val2) {
		res = 1
	}

	pos := in.write(p, 3, int64(res))
	p.pointer += in.len()

	if p.config.Print {
		fmt.Printf("Program %d: COMP: *%v = %v\n", p.id, pos, res)
	}
}

func (p *Program) rebase(in instruction) {
	val1 := in.read(p, 1)
	p.relBase += val1
	p.pointer += in.len()

	if p.config.Print {
		fmt.Printf("Program %d: RB: += %v // =%v\n", p.id, val1, p.relBase)
	}
}

func (p *Program) halt(value int64) {
	p.Exit <- value
	if p.config.Print {
		fmt.Printf("Program %d: Halt: Exiting program.\n", p.id)
	}
}
