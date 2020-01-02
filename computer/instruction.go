package computer

import (
	"strconv"
)

type parameter struct {
	num int
	mode int
}

func (p parameter) read(pm *Program) (value int64) {
	value = pm.memory[pm.pointer+int64(p.num)]

	switch p.mode {
	case 0: // position mode
		value = pm.memory[value]
	case 1: // immediate mode
		// use literal value
	case 2: // relative mode
		value = pm.memory[pm.relBase+value]
	}

	return value
}

func (p parameter) write(pm *Program, value int64) (addr int64) {
	addr = pm.memory[pm.pointer+int64(p.num)]

	switch p.mode {
	case 0: // position mode
		pm.memory[addr] = value
	case 1: // immediate mode
		panic("Cannot write to immediate mode parameter")
	case 2: // relative mode
		pm.memory[pm.relBase+addr] = value
	}

	return addr
}

type instruction struct {
	opcode int64
	params []parameter
}

func newInstruction(code int64) instruction {
	var instr instruction
	instr.opcode = code % 100

	switch instr.opcode {
	case 3, 4, 9:
		instr.params = make([]parameter, 1)
		instr.params[0] = parameter{1, int((code / 100) % 10)}
	case 5, 6:
		instr.params = make([]parameter, 2)
		instr.params[0] = parameter{1, int((code / 100) % 10)}
		instr.params[1] = parameter{2, int((code / 1000) % 10)}
	case 1, 2, 7, 8:
		instr.params = make([]parameter, 3)
		instr.params[0] = parameter{1, int((code / 100) % 10)}
		instr.params[1] = parameter{2, int((code / 1000) % 10)}
		instr.params[2] = parameter{3, int((code / 10000) % 10)}
	case 99:
		instr.params = make([]parameter, 0)
	default:
		panic("Unknown operation.")
	}

	return instr
}

func (in instruction) len() int64 {
	return int64(len(in.params) + 1)
}

func (in instruction) param(num int) parameter {
	if len(in.params) > (num-1) {
		return in.params[num-1]
	}

	panic("Instruction does not contain param " + strconv.Itoa(num))
}

func (in instruction) read(p *Program, param int) int64 {
	return in.param(param).read(p)
}

func (in instruction) write(p *Program, param int, value int64) int64 {
	return in.param(param).write(p, value)
}
