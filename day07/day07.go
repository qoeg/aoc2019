package day07

import (
	"sort"
	"strconv"
	"github.com/qoeg/aoc2019/computer"
)

func amplify(intcode, sequence []int, signal int) int {
	for i := range sequence {
		c := computer.New(i, intcode)
		c.Input <- sequence[i]
		c.Input <- signal

		go c.Run(false)

		signal = <-c.Output
	}

	return signal
}

func amplifyLoop(intcode, sequence []int, signal int) int {
	var amplifiers = make(map[int]*computer.Program, 5)

	for i := range sequence {
		p := computer.New(i, intcode)
		p.Input <- sequence[i]
		go p.Run(false)
		amplifiers[i] = p
	}

	return func() int {
		amplifiers[0].Input <- signal
		for {
			select {
			case signal = <-amplifiers[0].Output:
				amplifiers[1].Input <- signal
			case signal = <-amplifiers[1].Output:
				amplifiers[2].Input <- signal
			case signal = <-amplifiers[2].Output:
				amplifiers[3].Input <- signal
			case signal = <-amplifiers[3].Output:
				amplifiers[4].Input <- signal
			case signal = <-amplifiers[4].Output:
				amplifiers[0].Input <- signal
			case <-amplifiers[0].Exit:
			case <-amplifiers[1].Exit:
			case <-amplifiers[2].Exit:
			case <-amplifiers[3].Exit:
			case <-amplifiers[4].Exit:
				return signal
			}
		}
	}()
}

func permutations(set []int) [][]int {
	result := [][]int{}

	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			result = append(result, append([]int{}, a...))
			return
		}

		for i := k; i < len(set); i++ {
			a[k], a[i] = a[i], a[k]
			rc(a, k+1)
			a[k], a[i] = a[i], a[k]
		}
	}
	rc(set, 0)
	
	return result
}

// Answer1 returns the first puzzle answer
func Answer1() string {
	sequences := permutations([]int{0,1,2,3,4})

	signals := make([]int, len(sequences))
	for i, seq := range sequences {
		signals[i] = amplify(Input, seq, 0)
	}

	sort.Ints(signals)

	return strconv.Itoa(signals[len(signals)-1])
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	sequences := permutations([]int{5,6,7,8,9})

	signals := make([]int, len(sequences))
	for i, seq := range sequences {
		signals[i] = amplifyLoop(Input, seq, 0)
	}

	sort.Ints(signals)

	return strconv.Itoa(signals[len(signals)-1])
}
