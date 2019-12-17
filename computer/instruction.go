package computer

type instruction struct {
	program *Program
	opcode int64
	parameters int
	parameterModes []int
}

func (in instruction) read(num int) (value int64) {
	value = in.program.memory[in.program.pointer+int64(num)]

	switch in.parameterModes[num-1] {
	case 0: // position mode
		value = in.program.memory[value]
	case 1: // immediate mode
		// use literal value
	case 2: // relative mode
		value = in.program.memory[in.program.relBase+value]
	}

	return value
}

func (in instruction) write(num int, value int64) (position int64) {
	position = in.program.memory[in.program.pointer+int64(num)]

	switch in.parameterModes[num-1] {
	case 0: // position mode
		in.program.memory[position] = value
	case 1: // immediate mode
		panic("Cannot write to immediate mode parameter")
	case 2: // relative mode
		position = in.program.relBase+position
		in.program.memory[position] = value
	}

	return position
}
