package day04

import (
	"sort"
	"strconv"
)

func digits6(num int) []int {
	result := make([]int, 6) 
	result[0] = (num / 100000) % 10
	result[1] = (num / 10000) % 10
	result[2] = (num / 1000) % 10
	result[3] = (num / 100) % 10
	result[4] = (num / 10) % 10
	result[5] = num % 10
	return result
}

func equal(digits, comp []int) bool {
	for i, d := range digits {
		if d != comp[i] {
			return false
		}
	}
	return true
}

func escalating(digits []int) bool {
	sorted := make([]int, len(digits))
	copy(sorted, digits)
	sort.IntSlice(sorted).Sort()

	return equal(digits, sorted)
}

func pairs(digits []int) bool {
	last := -1
	for _, d := range digits {
		if d == last {
			return true
		}
		last = d
	}
	return false
}

func pairsOnly(digits []int) bool {
	count, last := 0, -1
	for i, d := range digits {
		if d == last {
			count++
			if count == 2 {
				if i+1 == len(digits) || digits[i+1] != d {
					return true
				}
			}
		} else {
			count = 1
		}

		last = d
	}
	return false
}

// Answer1 returns the first puzzle answer
func Answer1() string {
	numStart, _ := strconv.Atoi(Input[:6])
	numEnd, _ := strconv.Atoi(Input[7:])

	count := 0
	for n := numStart; n <= numEnd; n++ {
		digits := digits6(n)
		if escalating(digits) && pairs(digits) {
			count++
		}
	}

	return strconv.Itoa(count)
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	numStart, _ := strconv.Atoi(Input[:6])
	numEnd, _ := strconv.Atoi(Input[7:])

	count := 0
	for n := numStart; n <= numEnd; n++ {
		digits := digits6(n)
		if escalating(digits) && pairsOnly(digits) {
			count++
		}
	}

	return strconv.Itoa(count)
}
