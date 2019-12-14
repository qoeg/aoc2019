package day12

import (
	"strconv"
)

var test1 = []*moon{
	&moon{posx:-1, posy:0, posz:2},
	&moon{posx:2, posy:-10, posz:-7},
	&moon{posx:4, posy:-8, posz:8},
	&moon{posx:3, posy:5, posz:-1},
}

var test2 = []*moon{
	&moon{posx:-8, posy:-10, posz:0},
	&moon{posx:5, posy:5, posz:10},
	&moon{posx:2, posy:-7, posz:3},
	&moon{posx:9, posy:-8, posz:-3},
}

func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int64, integers ...int64) int64 {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func match(moons, state cluster, predicate func(cluster, cluster) bool, out chan int64) {
	var index int64
	for {
		moons.step()
		index++

		if predicate(moons, state) {
			out <- index
			return
		}
	}
}

// Answer1 returns the first puzzle answer
func Answer1() string {
	moons := cluster(Input)

	for i := 0; i < 1000; i++ {
		moons.step()
	}

	return strconv.Itoa(moons.energy())
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	moons := cluster(Input)
	xMatch := make(chan int64)
	yMatch := make(chan int64)
	zMatch := make(chan int64)

	go match(moons.clone(), moons, sameX, xMatch)
	go match(moons.clone(), moons, sameY, yMatch)
	go match(moons.clone(), moons, sameZ, zMatch)

	var xCount, yCount, zCount int64
	done := func() bool {
		return xCount > 0 && yCount > 0 && zCount > 0
	}
	
	for !done() {
		select {
		case xCount = <-xMatch:
		case yCount = <-yMatch:
		case zCount = <-zMatch:
		}
	}

	return strconv.FormatInt(lcm(xCount, yCount, zCount), 10)
}
