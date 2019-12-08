package computer

type instruction struct {
	pointer int
	opcode int
	parameters int
	parameterModes []int
}

func newInstruction(value, pointer int) instruction {
	var parameters int
	var parameterModes []int
	
	opcode := value % 100

	switch opcode {
	case 3, 4:
		parameters = 1
		parameterModes = make([]int, parameters)
		parameterModes[0] = (value / 100) % 10
	case 5, 6:
		parameters = 2
		parameterModes = make([]int, parameters)
		parameterModes[0] = (value / 100) % 10
		parameterModes[1] = (value / 1000) % 10
	case 1, 2, 7, 8:
		parameters = 3
		parameterModes = make([]int, parameters)
		parameterModes[0] = (value / 100) % 10
		parameterModes[1] = (value / 1000) % 10
		parameterModes[2] = (value / 10000) % 10
	case 99:
		parameters = 0
		parameterModes = make([]int, 0)
	default:
		panic("Unknown operation.")
	}

	return instruction{
		pointer,
		opcode,
		parameters,
		parameterModes,
	}
}

func (in instruction) parameter(num int, intcode []int) (value int) {
	position := intcode[in.pointer+num]

	value = position
	if in.parameterModes[num-1] == 0 {
		value = intcode[position]
	}
	return value
}
