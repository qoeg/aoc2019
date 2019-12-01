package day01

import (
	"strconv"
)

func baseFuel(mass int) (result int) {
	result = (mass / 3) - 2

	if result < 0 {
		return 0
	}

	return result
}

func compoundedFuel(mass int) (result int) {
	last := mass

	for {
		last = baseFuel(last)
		result += last

		if last == 0 {
			break
		}
	}

	return result
}

// Answer1 returns the first puzzle answer
func Answer1() string {
	sum := 0
	for _, v := range Input {
		sum += baseFuel(v)
	}

	return strconv.Itoa(sum)
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	sum := 0
	for _, v := range Input {
		sum += compoundedFuel(v)
	}
	
	return strconv.Itoa(sum)
}
