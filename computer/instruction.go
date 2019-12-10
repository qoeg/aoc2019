package computer

type instruction struct {
	program *Program
	opcode int64
	parameters int
	parameterModes []int
}

func newInstruction(program *Program, value int64) instruction {
	var parameters int
	var parameterModes []int
	
	opcode := value % 100

	switch opcode {
	case 3, 4, 9:
		parameters = 1
		parameterModes = make([]int, parameters)
		parameterModes[0] = int((value / 100) % 10)
	case 5, 6:
		parameters = 2
		parameterModes = make([]int, parameters)
		parameterModes[0] = int((value / 100) % 10)
		parameterModes[1] = int((value / 1000) % 10)
	case 1, 2, 7, 8:
		parameters = 3
		parameterModes = make([]int, parameters)
		parameterModes[0] = int((value / 100) % 10)
		parameterModes[1] = int((value / 1000) % 10)
		parameterModes[2] = int((value / 10000) % 10)
	case 99:
		parameters = 0
		parameterModes = make([]int, 0)
	default:
		panic("Unknown operation.")
	}

	return instruction{
		program,
		opcode,
		parameters,
		parameterModes,
	}
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
