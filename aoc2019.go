package main

import (
	"sort"
	"flag"

	"github.com/qoeg/aoc2019/day01"
	"github.com/qoeg/aoc2019/day02"
	"github.com/qoeg/aoc2019/day03"
	"github.com/qoeg/aoc2019/day04"
	"github.com/qoeg/aoc2019/day05"
	"github.com/qoeg/aoc2019/day06"
	"github.com/qoeg/aoc2019/day07"
	"github.com/qoeg/aoc2019/day08"
	"github.com/qoeg/aoc2019/day09"
	"github.com/qoeg/aoc2019/day10"
	"github.com/qoeg/aoc2019/day11"
	"github.com/qoeg/aoc2019/day12"
	"github.com/qoeg/aoc2019/day13"
)

var days = map[int]day{
	1: {1, day01.Answer1, day01.Answer2},
	2: {2, day02.Answer1, day02.Answer2},
	3: {3, day03.Answer1, day03.Answer2},
	4: {4, day04.Answer1, day04.Answer2},
	5: {5, day05.Answer1, day05.Answer2},
	6: {6, day06.Answer1, day06.Answer2},
	7: {7, day07.Answer1, day07.Answer2},
	8: {8, day08.Answer1, day08.Answer2},
	9: {9, day09.Answer1, day09.Answer2},
	10: {10, day10.Answer1, day10.Answer2},
	11: {11, day11.Answer1, day11.Answer2},
	12: {12, day12.Answer1, day12.Answer2},
	13: {13, day13.Answer1, day13.Answer2},
}

func main() {
	var all bool
	flag.BoolVar(&all, "all", false, "Setting to true will run all days")
	var requested dayFlags
	flag.Var(&requested, "days", "A list of days to run")
    flag.Parse()

	if all {
		for _, d := range sortedDays(days) {
			days[d].run()
		}
		return
	}

	if len(requested) > 0 {
		for _, d := range []day(requested) {
			d.run()
		}
		return
	}
	
	days[len(days)].run()
}

func sortedDays(days map[int]day) []int {
	keys := make([]int, len(days))
	i := 0
	for k := range days {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}
