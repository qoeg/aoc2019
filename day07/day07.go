package day07

import (
	"sort"
	"strconv"
	"github.com/qoeg/aoc2019/computer"
	"github.com/qoeg/aoc2019/util"
)

func amplify(intcode, sequence []int, signal int64) int64 {
	for i := range sequence {
		c := computer.New(i, computer.Config{})
		c.Input <- int64(sequence[i])
		c.Input <- int64(signal)

		go c.Run(intcode)

		signal = <-c.Output
	}

	return signal
}

func amplifyLoop(intcode, sequence []int, signal int64) int64 {
	var amplifiers = make(map[int]*computer.Program, 5)

	for i := range sequence {
		p := computer.New(i, computer.Config{})
		p.Input <- int64(sequence[i])
		go p.Run(intcode)
		amplifiers[i] = p
	}

	return func() int64 {
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

// Answer1 returns the first puzzle answer
func Answer1() string {
	sequences := util.Permutations([]int{0,1,2,3,4})

	signals := make([]int, len(sequences))
	for i, seq := range sequences {
		signals[i] = int(amplify(Input, seq, 0))
	}

	sort.Ints(signals)

	return strconv.Itoa(signals[len(signals)-1])
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	sequences := util.Permutations([]int{5,6,7,8,9})

	signals := make([]int, len(sequences))
	for i, seq := range sequences {
		signals[i] = int(amplifyLoop(Input, seq, 0))
	}

	sort.Ints(signals)

	return strconv.Itoa(signals[len(signals)-1])
}
