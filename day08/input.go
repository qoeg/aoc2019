package day08

// Input is the puzzle input
var Input = []int{2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,1,2,2,0,2,2,0,2,2,2,2,2,2,2,2,1,2,0,2,0,2,2,2,2,2,0,2,2,0,2,2,2,0,2,0,2,2,2,2,2,0,2,0,2,2,0,2,2,0,2,2,2,2,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,2,1,0,2,1,2,2,2,2,2,2,1,2,2,2,0,2,2,2,1,2,2,0,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,0,0,2,1,2,2,2,2,2,0,2,2,2,2,2,2,1,2,0,2,2,2,2,2,1,2,1,2,2,0,2,2,0,2,2,0,2,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,2,0,2,2,0,2,2,2,2,2,2,0,2,2,2,1,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,0,2,2,0,2,2,1,2,2,2,2,2,2,2,2,0,1,2,2,0,2,2,2,2,2,2,2,2,0,2,2,2,0,2,0,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,2,2,0,2,0,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,1,2,2,2,2,0,2,2,2,2,2,2,2,2,1,2,1,2,2,2,2,2,2,1,2,2,2,1,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,0,2,2,2,1,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,1,2,2,1,2,2,0,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,2,2,0,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,0,2,2,2,2,2,1,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,1,2,2,2,1,2,2,2,2,2,2,2,0,2,1,2,2,1,2,2,2,2,2,1,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,1,2,2,2,1,2,2,2,0,2,2,0,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,1,2,2,0,2,2,2,2,2,0,2,2,0,0,0,2,2,2,2,2,2,2,2,2,2,0,2,2,2,1,2,2,2,2,2,2,2,1,2,1,2,2,1,2,2,1,2,2,0,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,1,2,2,2,2,2,0,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,0,2,2,0,2,2,2,2,2,2,2,2,0,2,2,1,1,1,2,1,2,2,2,2,2,0,2,2,0,2,2,2,0,2,0,2,2,2,0,2,2,2,0,2,2,0,2,2,0,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,2,2,1,1,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,0,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,1,2,2,0,2,2,2,2,2,0,1,1,2,1,2,2,2,2,2,1,2,2,1,2,2,2,0,2,0,1,2,2,1,2,1,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,2,2,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,1,2,2,0,2,2,1,2,2,0,2,2,1,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,2,1,2,2,1,2,2,1,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,0,2,2,1,2,2,2,2,2,1,2,2,1,2,2,2,2,2,1,2,2,2,0,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,0,0,2,1,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,1,2,2,1,2,2,0,2,2,0,2,2,1,0,1,2,0,2,2,2,2,2,2,2,2,1,2,2,2,1,2,1,0,2,2,1,2,1,2,1,2,2,1,2,2,0,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,0,2,2,2,2,2,2,2,2,2,2,0,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,1,2,2,2,2,2,1,0,2,2,2,2,1,2,2,2,2,1,2,2,0,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,1,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,1,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,2,2,2,1,2,2,2,1,2,2,1,2,2,1,2,2,1,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,0,2,2,2,0,2,2,1,2,2,1,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,2,2,2,2,2,1,2,2,1,2,2,2,0,0,2,2,2,2,2,2,2,0,2,2,2,2,2,2,0,2,2,2,2,2,1,2,0,2,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,2,2,0,2,0,2,2,2,2,2,2,1,0,2,2,2,2,2,2,1,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,0,2,2,1,2,2,2,2,2,0,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,0,0,2,2,2,0,2,0,2,2,2,2,2,2,0,0,2,2,0,2,2,2,0,2,2,0,2,2,0,2,2,2,2,0,2,2,2,2,1,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,1,2,2,1,2,2,2,2,2,0,2,2,1,1,0,2,1,2,2,2,2,2,1,2,2,1,2,2,2,0,2,2,0,2,2,0,2,0,2,1,2,2,2,2,2,1,2,2,1,2,2,2,2,1,2,2,2,2,1,2,2,2,1,1,2,2,0,2,2,2,2,2,2,2,2,2,1,0,2,2,1,2,2,2,0,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,1,2,2,2,2,2,2,2,2,0,2,2,1,0,1,2,2,2,2,2,2,2,0,2,2,0,2,2,2,2,2,0,1,2,2,0,2,1,2,1,2,2,0,2,2,1,2,2,1,2,2,2,2,0,2,2,2,2,0,2,2,2,0,2,2,2,1,1,2,0,2,2,2,2,2,2,0,2,2,2,1,2,2,2,1,2,2,1,2,2,2,2,2,2,2,1,2,2,2,2,0,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,2,2,0,2,2,2,0,2,0,2,2,2,1,2,0,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,0,2,2,2,0,1,2,2,2,2,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,2,2,0,2,2,2,2,2,1,2,2,1,2,1,2,0,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,1,2,2,0,2,1,2,2,2,2,1,2,2,0,2,2,0,2,2,2,2,1,1,2,2,0,1,2,2,2,2,1,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,1,2,2,1,2,2,2,2,2,0,2,1,2,2,2,2,2,2,2,1,2,2,0,2,2,2,0,2,1,1,2,2,2,2,0,2,1,2,2,1,2,2,1,2,2,2,2,2,2,2,1,0,2,2,1,0,2,1,2,2,0,2,2,1,0,2,0,2,2,2,2,2,2,1,1,2,2,1,2,2,2,0,2,2,1,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,0,2,0,2,2,2,2,2,1,2,2,2,2,2,2,1,2,0,1,2,2,1,2,1,2,0,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,1,0,2,2,0,2,2,2,0,2,2,1,2,2,0,2,0,2,2,0,2,2,2,2,1,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,1,2,2,1,2,2,1,2,2,1,1,0,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,2,0,2,1,2,2,2,2,1,2,2,2,2,2,0,2,2,2,2,0,2,2,1,2,0,2,0,2,0,2,2,2,1,2,2,2,2,2,2,2,2,2,2,0,2,2,0,2,2,2,0,2,2,0,2,2,0,2,0,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,1,2,2,1,2,2,1,2,2,0,2,2,1,1,2,2,1,2,2,2,2,2,1,2,2,0,2,2,2,1,2,0,1,2,2,1,2,1,2,0,2,2,2,2,2,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,1,2,2,2,2,2,1,0,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,0,2,2,2,2,2,1,0,1,2,1,2,2,2,2,2,0,2,2,2,2,2,2,2,2,1,1,2,2,1,2,1,2,0,2,2,1,2,2,0,2,2,0,2,2,2,2,2,0,2,0,0,0,2,1,2,1,2,2,2,2,0,2,0,2,2,2,2,2,2,0,1,2,2,1,2,2,2,1,2,2,0,2,2,2,2,0,2,2,0,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,1,2,2,2,2,2,2,2,2,0,2,2,2,0,0,2,1,2,2,2,2,2,0,2,2,0,2,2,2,1,2,1,1,2,2,1,2,2,2,0,2,2,0,2,2,0,2,2,1,2,2,2,2,0,1,2,0,1,0,2,0,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,0,2,2,2,2,2,0,2,2,2,2,2,1,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,1,2,1,1,2,2,1,2,1,2,0,2,2,2,2,2,0,2,2,2,2,2,2,2,0,2,2,1,2,0,2,1,2,1,0,2,2,0,2,2,2,2,2,2,2,2,2,1,1,2,2,0,2,2,2,1,2,2,2,2,2,0,2,0,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,2,2,2,1,2,2,0,2,2,2,2,2,0,0,0,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,1,2,1,2,2,0,2,2,0,2,2,2,2,2,2,2,2,2,2,0,2,1,2,1,2,0,0,2,2,2,0,2,1,2,2,2,2,2,2,0,1,2,2,0,2,2,2,2,2,2,1,2,2,0,2,0,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,1,2,2,2,2,2,1,2,2,0,2,2,1,2,2,2,1,0,2,2,2,2,2,2,2,1,2,2,0,2,2,2,1,2,0,2,2,2,2,2,1,2,2,2,2,1,2,2,1,2,2,2,2,2,2,2,2,1,2,2,2,1,2,0,2,2,1,2,2,0,0,2,0,2,2,2,2,2,2,0,0,2,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,1,2,2,0,2,2,0,2,2,1,2,2,0,2,2,1,0,0,2,1,2,2,2,2,2,2,2,2,2,2,2,2,1,2,0,0,2,2,0,2,0,2,1,2,2,0,2,2,0,2,2,1,2,2,2,2,1,0,2,0,0,1,2,0,2,1,0,2,2,0,1,2,0,2,2,2,2,2,2,2,0,2,2,0,2,2,2,2,2,2,1,2,2,1,2,2,2,2,1,2,2,2,1,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,1,1,2,2,0,2,2,2,2,2,0,2,2,0,2,2,2,1,1,2,1,2,2,2,2,2,2,2,2,1,2,2,2,0,2,0,1,2,2,0,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,0,2,0,2,0,2,1,2,2,0,2,2,0,2,2,0,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,1,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,0,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,0,2,0,2,2,2,2,2,2,0,2,2,0,2,2,0,2,2,2,0,1,2,0,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,0,2,2,1,2,2,2,2,1,0,2,1,2,1,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,2,0,2,2,1,2,2,2,2,1,2,2,0,2,2,2,0,2,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,2,0,1,2,2,0,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,0,2,2,2,2,2,1,2,2,1,2,2,2,2,2,0,1,2,2,2,2,1,2,0,2,2,1,2,2,0,2,2,2,2,2,2,2,1,2,2,1,1,2,2,0,2,0,2,2,2,1,1,2,2,2,2,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,1,2,2,1,2,2,2,2,2,1,2,2,0,2,2,1,0,2,2,1,2,2,2,2,2,1,2,2,2,2,2,2,2,2,1,2,2,2,0,2,1,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,1,1,1,2,1,2,0,0,2,2,2,1,2,0,2,2,2,2,2,0,0,1,2,2,0,2,2,2,2,2,2,1,2,2,1,2,1,2,2,1,2,2,2,1,0,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,0,2,2,2,2,1,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,0,2,0,2,2,2,1,2,2,1,2,0,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,1,1,2,2,2,2,1,2,0,2,2,0,2,2,2,2,2,0,2,2,2,2,0,1,2,2,2,2,2,2,2,0,2,2,2,2,1,2,0,2,2,2,2,2,1,1,0,2,2,1,2,2,2,1,2,2,2,2,2,2,2,1,2,2,1,2,2,2,2,0,2,2,2,2,2,2,0,1,2,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,1,2,2,2,2,2,1,2,0,0,2,2,2,2,2,2,2,1,1,2,2,0,2,2,2,1,0,2,0,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,2,0,2,2,0,2,2,1,2,2,0,2,2,2,2,1,0,2,1,0,2,2,1,2,1,2,2,2,1,1,2,1,2,2,2,2,2,0,0,0,2,2,0,2,2,2,1,2,2,1,2,2,2,2,2,2,2,0,2,2,2,1,0,2,2,2,2,2,2,1,0,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,0,2,2,2,2,2,0,2,2,0,2,1,2,2,2,2,2,2,2,1,2,2,0,2,2,2,2,2,1,2,2,2,2,2,2,2,2,0,2,2,2,1,2,0,2,0,2,2,0,2,2,1,2,2,0,2,2,2,2,2,1,2,0,1,0,2,1,2,1,0,2,2,0,1,2,2,2,2,2,2,2,0,0,2,2,2,1,2,2,2,0,2,2,0,2,2,0,2,0,2,2,0,2,2,2,0,1,2,2,2,2,2,2,1,1,1,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,0,1,2,2,2,2,1,2,0,2,2,2,2,2,2,0,2,2,2,2,2,1,2,2,2,0,2,2,0,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,1,2,2,1,2,1,2,2,2,2,0,2,2,2,2,2,1,2,2,2,2,2,0,2,2,0,2,2,0,2,2,1,2,2,2,1,2,2,2,2,2,2,2,1,0,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,0,2,2,1,2,2,2,2,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,0,0,2,2,2,2,2,2,0,2,2,1,2,2,0,2,0,1,0,2,0,2,2,0,0,2,2,1,2,2,2,2,2,1,2,2,1,2,2,2,2,2,0,1,2,2,2,2,1,2,0,2,2,2,2,2,0,2,2,2,2,2,2,2,0,2,2,2,1,1,2,0,2,0,0,2,2,0,0,2,1,2,2,2,2,2,1,0,2,2,2,2,2,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,1,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,1,2,2,1,2,2,0,2,2,1,0,1,0,2,2,1,1,2,2,1,2,2,2,2,2,0,2,2,1,2,2,2,2,2,0,0,2,2,1,2,0,2,1,2,2,1,2,2,1,2,2,1,2,2,2,2,1,0,2,1,0,1,2,1,2,1,2,2,2,0,0,2,1,2,2,2,2,2,1,2,2,2,2,0,2,2,2,2,2,2,2,2,2,1,2,0,2,2,0,2,2,2,1,1,2,2,2,2,2,1,1,1,0,2,2,2,2,2,0,2,2,2,2,2,2,2,0,2,2,2,2,0,1,2,2,2,2,2,2,2,2,1,0,0,2,2,2,1,2,2,0,2,2,2,2,2,1,2,2,0,0,0,0,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,1,2,2,2,1,2,0,0,2,2,2,2,2,2,0,2,2,1,2,2,0,2,2,0,2,2,2,2,0,2,2,1,1,0,2,0,2,2,2,2,2,1,0,2,2,2,0,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,1,0,2,1,2,2,2,2,1,2,2,2,2,1,2,2,2,2,2,1,1,2,2,2,2,1,2,2,1,2,2,1,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,1,1,2,2,2,0,2,2,2,2,2,1,2,2,2,2,1,0,2,0,2,2,2,1,1,1,2,1,2,2,2,2,2,1,2,2,1,2,2,2,0,2,2,1,2,2,1,2,2,2,2,2,2,2,2,2,1,2,1,1,2,2,2,2,0,1,2,1,2,2,2,1,2,0,0,2,2,0,2,2,2,2,0,2,2,2,1,0,0,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,0,1,0,1,2,2,2,2,2,1,2,2,0,2,2,2,2,1,2,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,1,2,2,2,1,2,1,2,2,2,0,2,2,1,2,1,0,2,2,2,2,2,2,0,1,2,2,2,2,2,2,2,0,2,2,0,2,2,2,1,2,0,1,2,2,1,2,0,2,2,2,2,1,2,2,0,2,0,0,2,2,2,2,2,0,2,1,1,1,2,0,2,2,2,2,2,0,0,2,2,2,0,2,2,2,0,0,2,2,2,0,2,2,2,1,2,2,2,2,2,1,2,2,2,2,1,2,2,2,1,1,2,2,2,2,2,1,0,1,1,2,2,2,2,2,0,2,2,0,2,2,2,2,2,2,2,2,2,1,1,2,2,2,2,2,2,2,2,2,1,1,2,2,2,1,2,0,1,2,2,2,2,2,1,2,2,2,1,0,0,2,2,0,0,2,2,1,2,2,2,2,2,1,2,2,0,2,2,2,2,2,2,0,2,2,2,2,0,2,0,2,2,1,2,2,2,2,1,2,2,2,2,2,1,1,2,0,1,2,2,2,2,1,2,2,2,0,0,2,0,2,2,2,2,2,2,1,1,2,2,0,2,2,2,2,2,2,1,0,2,1,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,0,2,1,2,2,0,2,2,1,2,2,2,0,2,2,2,2,2,2,2,2,0,0,2,2,2,2,2,2,2,2,2,2,0,2,2,2,0,2,1,1,2,2,0,2,1,0,2,2,1,0,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,0,2,2,1,2,2,2,2,0,2,2,2,1,1,2,2,0,2,0,0,2,2,2,2,2,2,2,2,0,1,2,0,2,2,1,2,2,1,1,2,2,2,0,2,2,2,0,0,1,2,2,2,2,2,2,0,2,2,1,1,2,1,2,0,2,2,1,2,2,2,1,1,2,2,2,2,2,0,0,2,0,2,2,0,2,2,2,2,2,2,0,2,2,2,0,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,1,2,0,1,2,2,0,2,1,0,2,1,1,2,1,1,2,2,2,2,0,2,1,2,2,2,2,2,0,2,2,2,2,2,2,2,2,1,2,2,2,2,2,1,2,1,2,0,0,2,2,1,2,2,0,2,2,2,2,2,2,2,1,0,1,2,0,2,1,0,2,2,1,1,2,2,2,1,2,2,2,0,1,0,2,2,0,2,2,2,1,2,2,0,0,2,2,2,0,2,2,0,2,2,2,1,1,2,2,2,2,2,2,1,0,0,2,2,0,2,2,0,2,2,2,0,2,2,2,0,2,2,2,2,0,1,2,2,2,1,2,2,2,2,2,2,0,2,2,1,0,2,0,2,2,2,2,2,1,2,2,0,0,2,0,0,2,2,2,1,1,2,1,2,2,2,2,2,1,2,2,0,2,2,2,2,2,0,2,2,2,1,2,2,2,2,2,0,1,2,2,2,2,0,0,2,2,2,2,2,1,2,0,2,0,2,0,2,0,0,2,2,0,1,2,0,2,1,2,2,2,0,1,1,2,2,2,2,2,2,0,2,2,1,2,2,0,2,2,2,2,2,2,2,2,1,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,0,2,2,2,2,2,1,2,2,2,0,2,2,2,2,1,0,0,2,2,2,2,2,1,2,2,2,1,2,1,0,2,0,1,1,0,2,2,2,0,1,0,2,2,2,2,2,2,2,0,0,2,1,2,2,2,1,2,2,0,2,2,2,2,2,2,0,0,2,0,2,2,0,2,1,2,2,2,2,2,0,0,2,0,1,1,2,2,2,0,2,2,2,0,1,2,2,2,2,2,2,2,0,2,0,2,2,2,2,2,2,2,2,2,2,1,2,0,2,2,2,2,0,2,2,2,1,0,2,2,2,2,2,2,0,1,0,2,2,2,2,2,0,2,2,0,0,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,2,2,1,0,2,2,2,2,1,2,2,2,1,2,2,0,2,2,2,2,0,2,0,1,0,2,2,0,0,1,2,0,2,2,2,2,2,1,2,2,1,2,2,2,1,2,1,2,2,2,0,2,0,2,0,1,1,2,2,2,1,2,1,0,2,2,2,2,0,1,2,0,0,1,2,2,2,0,1,2,2,2,1,2,0,2,1,2,2,2,0,2,1,0,2,0,2,2,2,2,2,2,1,0,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,1,2,2,1,2,2,1,2,2,1,1,2,2,2,0,2,2,2,2,1,0,2,2,2,0,2,2,2,2,1,0,1,2,2,2,2,2,2,1,2,2,0,2,0,1,2,1,0,1,0,1,2,2,2,1,1,2,1,2,2,2,2,2,2,0,2,1,2,2,2,0,2,2,1,2,2,2,2,0,2,1,2,1,0,2,2,0,2,1,1,2,2,2,2,2,1,2,1,0,1,2,0,2,2,1,2,2,2,2,2,2,2,1,2,2,2,1,1,0,2,2,0,2,2,2,1,2,2,2,1,2,1,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,0,1,1,2,2,2,0,2,2,2,2,2,0,2,2,2,2,1,2,2,2,2,0,1,2,2,2,2,2,2,2,0,1,0,1,2,2,2,1,2,1,2,2,2,0,2,1,0,2,1,1,1,1,0,2,2,0,0,1,2,2,2,2,2,2,2,1,2,2,1,2,2,2,0,2,1,1,2,2,2,2,1,2,2,1,0,0,2,2,2,2,1,0,2,2,2,2,1,1,2,0,0,2,2,1,2,2,2,2,2,2,1,2,0,2,1,2,2,2,2,2,0,2,2,1,2,2,2,0,2,2,1,1,2,1,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,1,2,0,0,2,2,2,2,2,2,0,2,0,1,2,2,2,0,2,2,2,2,0,2,2,2,2,0,2,2,2,2,2,1,0,2,2,0,1,2,2,1,2,2,0,2,1,2,2,1,2,2,1,2,2,2,1,2,2,2,1,2,2,2,2,2,0,1,2,0,2,2,2,2,2,1,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,2,2,0,2,2,2,1,2,2,2,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,0,0,2,2,1,2,2,2,0,2,2,0,0,2,0,2,2,2,2,2,2,2,2,0,1,2,2,2,2,2,0,2,2,1,2,2,2,2,2,2,1,2,1,2,2,2,2,0,2,2,2,2,1,0,2,2,2,2,2,2,2,1,0,2,0,2,2,1,1,2,0,2,2,2,0,2,1,1,2,1,1,0,1,2,2,2,1,1,0,2,0,2,2,2,2,2,1,2,2,0,2,2,2,1,2,1,1,2,2,2,2,0,2,1,1,2,2,2,2,2,2,2,1,2,2,2,2,1,0,0,1,2,0,2,2,2,2,1,2,2,1,2,2,2,2,0,2,2,2,1,1,1,1,2,0,2,2,2,0,2,2,1,2,2,2,2,0,2,2,1,2,2,2,1,1,2,2,2,2,2,0,0,0,2,2,2,1,2,2,1,0,2,1,1,2,2,2,1,2,2,2,2,1,2,2,2,2,0,2,2,2,1,2,0,0,2,2,0,0,2,1,1,2,2,2,2,1,0,2,1,1,2,2,1,2,2,0,2,0,2,2,2,2,2,2,2,1,0,2,0,2,2,2,2,2,1,2,2,2,2,2,0,2,1,0,2,0,2,2,1,2,0,0,2,2,2,2,1,1,1,2,0,0,2,1,2,1,1,2,2,0,2,2,0,2,2,2,2,2,0,0,1,1,2,2,2,2,2,0,2,2,2,1,2,2,2,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,0,2,0,2,2,0,2,2,1,1,2,1,0,2,2,2,1,2,2,2,2,0,1,2,2,2,1,2,2,1,2,0,0,1,2,2,0,1,2,1,1,2,2,2,2,1,2,2,2,0,1,0,0,2,2,1,2,1,2,1,2,2,2,2,2,0,2,2,1,2,2,2,1,2,1,0,2,2,1,2,1,2,2,1,2,1,2,2,1,2,1,2,2,2,2,2,0,1,1,2,0,2,2,0,2,1,1,2,2,2,1,2,1,2,1,2,2,2,1,1,1,1,2,0,2,2,2,0,2,2,2,0,2,2,2,0,2,2,1,2,2,2,2,1,2,2,2,2,2,1,1,1,1,2,1,2,2,2,2,1,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,1,2,2,1,0,0,2,1,2,2,2,1,2,1,1,2,2,0,2,0,0,2,0,2,1,0,1,2,2,1,0,1,2,1,2,2,1,2,2,2,1,2,1,2,2,2,0,2,2,1,2,2,2,2,0,2,0,1,1,0,2,2,1,2,2,1,2,2,2,2,0,1,1,1,0,2,2,0,2,2,0,2,2,2,0,2,2,2,2,2,2,2,1,2,1,2,2,2,2,2,2,2,2,2,1,2,2,2,2,1,2,2,2,2,1,2,0,1,2,2,2,2,2,0,2,1,2,2,2,1,2,2,2,0,2,1,0,2,2,2,2,1,2,2,2,1,0,2,2,2,2,2,2,0,0,0,0,0,2,2,2,0,2,0,2,2,2,1,2,2,1,2,1,2,0,1,2,2,2,2,0,1,2,2,2,2,0,2,2,2,1,2,1,2,2,2,1,2,2,2,2,2,0,2,2,2,0,1,0,2,2,2,2,2,0,2,2,2,2,2,1,2,1,2,1,1,2,1,2,2,0,2,2,0,0,2,0,2,0,2,2,2,0,2,2,1,2,0,2,2,2,1,2,2,2,2,2,0,2,2,2,2,2,2,1,2,0,2,2,2,2,2,2,1,1,2,0,2,2,2,2,2,2,2,2,2,0,1,2,2,2,0,2,2,2,1,1,2,2,2,2,2,2,2,1,1,0,2,2,2,1,1,2,2,1,2,2,0,2,1,1,2,1,2,2,1,0,2,2,1,2,0,2,1,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,1,2,2,2,2,0,2,1,2,0,2,2,2,1,2,2,0,2,2,2,2,0,2,0,0,0,1,2,2,2,0,1,2,2,0,2,2,2,2,1,2,2,2,2,0,1,0,2,2,2,2,2,0,2,2,2,0,2,1,2,2,2,2,2,2,0,2,1,0,2,2,2,2,2,2,1,2,2,2,2,1,2,2,1,0,2,0,1,1,2,2,1,0,2,2,2,2,2,2,2,2,1,2,2,1,1,2,0,0,2,2,2,1,2,2,1,2,2,2,2,2,0,2,1,0,0,2,1,2,2,1,0,1,2,0,2,2,0,2,2,0,1,2,1,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,0,2,2,2,2,2,2,0,2,2,2,2,0,0,0,1,1,2,2,1,2,1,0,2,2,0,1,2,1,2,2,2,2,2,2,0,2,0,2,2,2,2,2,1,2,2,1,2,2,2,2,2,2,2,1,2,1,2,2,2,2,2,2,2,2,2,0,1,2,2,1,1,2,2,2,0,2,1,1,1,2,2,0,2,2,2,2,2,1,2,2,2,0,2,2,0,2,1,0,1,2,2,1,2,2,0,1,2,2,1,2,2,2,2,1,1,2,2,0,2,2,1,2,2,2,0,2,2,1,2,2,2,0,2,0,2,2,2,0,2,0,1,2,2,2,2,1,2,0,1,0,0,2,2,2,2,2,1,2,2,2,2,2,2,1,2,0,1,2,1,2,1,1,2,2,2,2,2,0,2,1,2,2,2,0,0,1,0,2,0,2,2,2,2,2,2,1,0,2,0,2,2,2,2,0,2,0,2,1,1,2,2,2,2,2,1,1,0,1,2,2,0,2,2,0,2,2,2,2,1,2,2,2,1,2,2,2,2,0,2,2,2,1,2,2,1,0,0,0,1,2,2,0,1,2,1,1,2,2,1,2,1,2,2,1,0,0,2,2,2,2,2,1,0,2,2,2,2,2,2,2,0,2,2,0,2,1,2,1,2,1,0,2,2,1,2,2,2,0,1,2,1,0,2,2,2,0,0,2,2,2,2,2,0,2,0,0,0,2,0,2,2,1,2,2,0,1,2,2,2,1,2,2,2,2,1,0,2,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,1,2,0,1,2,2,2,2,2,0,2,0,0,2,2,2,2,2,1,0,2,1,0,1,2,2,2,1,2,2,2,0,0,2,2,2,0,2,2,2,1,2,2,2,2,2,2,1,2,1,2,2,2,2,2,2,1,2,1,2,0,0,0,2,2,0,1,0,2,1,2,2,0,2,2,1,2,2,0,2,1,2,0,2,0,0,2,2,0,2,1,2,0,1,0,1,1,2,0,2,1,2,2,2,2,2,0,0,0,0,2,1,2,2,2,0,0,2,2,2,2,2,2,2,0,2,2,2,1,2,2,1,2,2,2,2,2,1,2,2,1,2,2,0,2,0,2,2,2,2,2,2,2,1,2,2,2,2,2,0,2,1,2,2,2,0,2,2,2,0,2,0,0,0,2,2,0,0,2,2,2,2,2,2,2,2,1,2,2,2,1,2,2,0,2,2,0,1,2,0,1,2,2,2,2,2,1,2,1,1,0,1,2,2,2,1,1,0,2,2,2,2,1,2,2,0,1,2,0,2,0,2,2,2,1,0,2,0,2,2,0,2,2,2,1,0,1,2,0,2,2,2,2,2,2,2,0,0,2,0,2,2,2,0,2,0,1,2,2,2,0,2,1,2,2,2,2,2,0,0,1,2,2,1,2,2,2,2,2,2,0,0,2,2,2,0,2,2,2,2,2,2,0,1,2,2,2,2,2,0,1,2,1,2,2,0,2,2,1,1,2,0,0,0,2,2,0,1,2,2,2,0,2,2,2,2,2,1,2,2,2,2,0,1,2,2,1,1,2,1,0,2,2,1,2,0,1,2,1,1,1,0,0,2,2,2,2,2,2,0,2,2,2,2,2,2,2,2,0,2,1,2,2,2,0,0,2,2,2,2,1,2,0,0,1,0,2,2,0,2,2,0,2,2,2,2,2,1,1,0,0,0,2,1,2,0,2,2,2,1,0,2,0,2,2,2,2,2,1,1,1,1,2,1,2,2,2,1,2,2,0,2,2,1,2,1,2,2,1,2,2,2,0,2,2,2,2,0,2,0,0,2,0,2,1,0,2,2,2,1,2,0,1,1,2,2,1,2,2,2,2,2,2,2,2,2,0,2,2,0,1,1,1,0,2,2,2,2,2,1,0,2,2,1,2,2,2,2,0,1,0,2,1,2,2,2,2,2,2,0,2,2,0,2,2,2,0,2,1,2,2,2,0,2,0,1,2,1,1,2,0,2,1,0,0,2,1,2,0,2,1,2,2,2,2,2,0,0,1,1,2,2,2,2,2,2,0,2,2,1,0,2,2,2,2,2,2,2,1,2,0,0,2,0,2,2,2,1,2,2,1,0,2,1,2,1,2,2,0,1,2,2,1,2,2,2,2,0,2,2,0,0,2,2,2,2,2,2,0,1,2,0,0,0,2,2,0,0,2,2,2,1,0,2,2,2,1,2,2,1,2,0,1,2,2,2,0,1,2,1,1,2,2,1,2,0,2,2,0,1,2,0,2,2,2,2,2,1,2,1,2,2,2,2,2,1,0,2,1,2,1,2,0,2,2,2,2,1,1,2,1,2,1,2,1,0,2,2,1,2,0,0,2,2,2,2,0,1,0,2,1,2,2,2,2,2,0,2,2,0,0,2,1,2,1,2,2,2,0,2,2,1,2,2,2,2,2,2,2,2,0,2,2,2,2,0,2,2,2,1,2,2,1,1,2,2,2,2,2,0,2,2,0,2,1,0,2,2,2,2,2,0,1,0,1,2,2,1,2,2,2,2,1,2,2,2,0,1,2,0,2,2,1,2,2,2,2,2,2,0,0,2,2,1,2,2,1,2,1,1,2,1,0,2,2,0,1,1,2,1,2,2,0,2,2,2,2,2,0,2,0,2,2,2,1,2,0,0,2,2,0,2,1,0,0,0,0,2,2,2,0,1,2,2,2,2,2,0,2,0,1,2,2,2,2,1,0,2,2,1,0,2,0,2,0,2,2,2,2,1,1,1,2,1,2,2,2,0,2,2,2,2,2,1,2,2,2,2,1,0,1,2,0,2,2,2,2,2,2,0,0,0,1,2,0,0,2,2,2,1,2,0,1,0,2,2,0,2,2,2,2,2,1,2,2,2,2,2,2,2,1,1,2,1,2,2,2,2,2,1,2,2,2,1,2,0,1,2,2,0,2,1,2,2,2,2,1,0,2,0,2,2,0,2,2,0,1,2,2,2,0,2,0,2,0,2,2,2,0,2,0,2,0,2,2,0,1,2,1,2,0,0,2,2,2,2,0,0,1,1,0,0,2,0,2,0,2,2,2,1,0,2,0,2,2,2,2,2,2,2,2,2,2,2,2,2,2,0,2,2,2,2,2,1,2,0,2,2,1,0,1,2,2,0,2,2,2,0,2,2,0,0,2,2,1,2,2,2,1,2,2,2,0,1,1,2,0,0,2,2,2,0,0,2,2,2,0,1,2,2,1,0,1,0,2,2,0,1,2,2,2,2,2,2,2,2,1,2,1,1,1,0,0,2,2,2,0,1,2,2,2,2,1,2,2,0,2,2,0,2,2,2,1,2,1,1,1,0,1,2,2,2,0,2,1,2,2,1,2,2,1,2,2,2,2,2,1,0,1,1,1,2,2,1,2,1,1,2,2,0,1,2,1,2,1,2,2,2,2,2,0,0,2,0,2,2,2,0,2,2,0,2,2,2,2,0,2,2,0,2,0,2,0,2,2,2,2,0,2,1,2,2,1,2,1,1,2,2,0,0,2,0,0,2,0,2,0,1,2,2,2,0,2,2,2,2,1,2,2,0,2,1,2,0,2,2,0,0,2,2,0,2,2,1,2,2,1,2,1,1,1,1,1,2,2,2,2,1,2,2,2,2,0,2,2,1,1,2,1,2,2,0,2,2,1,0,2,0,1,2,2,2,0,2,1,1,1,2,2,2,2,2,2,2,2,2,0,2,1,0,0,0,2,0,2,1,2,2,2,2,0,2,1,2,1,2,2,2,2,0,1,0,2,1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,1,0,0,2,0,1,2,2,2,1,2,2,1,2,0,2,2,1,2,2,0,0,0,1,0,2,0,2,2,2,2,2,2,2,0,2,2,2,1,2,2,0,2,0,1,0,2,2,0,1,2,0,1,2,2,2,2,1,2,2,0,0,2,1,2,2,2,0,0,1,2,1,2,2,2,2,2,0,1,2,0,2,0,2,2,2,1,1,2,0,1,2,2,2,2,0,2,0,2,2,0,2,1,1,2,2,2,2,2,1,1,1,2,1,2,2,2,0,2,2,2,2,0,2,0,2,1,2,2,2,0,0,0,0,2,0,2,2,2,0,2,2,1,2,2,1,2,1,2,2,1,1,1,2,2,2,2,2,2,0,2,1,0,1,0,2,0,2,2,2,1,2,2,0,0,1,0,2,2,0,2,2,2,1,1,2,2,2,1,0,2,1,2,2,1,1,2,2,1,0,0,1,2,2,2,2,2,2,1,2,2,0,1,1,0,2,2,1,2,0,2,1,2,2,2,2,2,0,0,2,0,2,1,0,2,2,0,0,1,0,0,2,2,2,0,2,1,0,2,1,0,2,1,2,2,2,2,2,1,0,0,2,0,2,2,2,2,1,2,2,2,2,0,2,2,2,2,2,2,2,1,2,2,0,2,1,2,2,2,2,2,2,1,1,2,0,2,1,2,2,0,0,1,2,2,2,2,2,2,2,2,2,0,0,2,2,2,1,2,2,0,1,2,1,1,2,1,2,0,2,2,2,2,0,1,2,2,2,1,0,2,2,2,1,1,0,2,2,0,0,2,2,2,2,2,1,2,1,2,2,1,0,0,1,2,2,2,0,1,2,2,1,2,2,2,2,2,2,2,2,2,2,2,1,2,2,1,1,2,1,1,2,2,2,2,1,0,0,0,0,0,2,0,2,2,2,2,2,1,0,1,0,1,2,0,0,2,2,0,2,2,0,0,2,2,2,1,2,2,2,0,1,2,1,2,1,2,2,2,0,2,2,0,0,2,2,1,1,2,2,1,2,1,2,0,2,2,2,2,0,2,1,2,0,1,2,2,1,2,2,2,2,2,0,1,1,1,2,0,2,2,2,2,0,1,2,2,2,2,1,2,2,2,0,1,2,2,2,0,2,0,1,2,2,2,0,2,1,2,2,0,0,1,1,2,2,2,0,2,0,2,0,2,2,0,2,2,2,1,2,2,2,0,1,2,2,0,0,1,2,2,2,1,2,2,2,2,0,1,2,0,2,0,1,2,2,2,2,0,1,0,0,0,1,2,1,2,0,2,2,2,1,2,2,2,2,0,2,0,2,1,0,1,0,2,2,2,2,0,1,2,2,0,2,2,0,0,0,2,2,0,2,0,2,2,2,2,2,2,1,2,1,2,0,1,2,2,2,2,2,2,2,1,1,1,0,1,2,1,2,2,2,2,2,1,2,2,2,0,1,2,2,1,1,0,1,2,2,0,2,2,0,1,2,2,0,2,2,0,2,2,2,1,0,1,2,2,2,0,1,2,1,2,2,1,2,2,0,1,2,0,2,1,2,1,2,1,1,1,1,1,2,2,2,2,0,0,1,2,1,1,2,1,2,2,2,2,2,2,0,0,0,0,0,2,1,2,0,0,2,2,1,0,2,0,2,2,1,1,2,1,0,2,2,2,1,2,2,0,2,2,2,1,0,2,2,1,0,2,2,1,2,0,2,1,2,2,2,2,2,2,1,0,1,2,2,0,2,2,2,1,0,1,1,1,1,1,2,2,1,2,2,2,1,1,2,2,2,2,1,2,2,1,1,1,2,2,2,0,0,1,2,0,2,2,1,2,0,0,2,0,0,1,0,1,2,2,1,2,0,2,0,2,2,0,2,2,0,0,2,1,2,0,2,1,2,0,1,1,0,0,2,1,2,1,0,2,2,2,0,0,2,0,2,2,2,2,2,1,0,0,0,0,1,1,1,2,2,1,2,2,2,2,2,1,2,1,0,2,2,2,2,1,1,2,0,2,2,1,0,2,2,0,0,2,1,2,2,2,2,1,0,2,2,0,0,2,2,2,2,2,2,0,1,2,2,0,2,2,2,1,0,1,0,0,1,0,2,2,1,1,2,2,2,1,2,2,2,0,1,2,1,0,2,0,0,2,2,2,1,1,0,2,2,2,1,2,2,0,2,0,2,0,2,1,2,2,1,1,0,2,2,2,2,1,2,2,2,2,2,0,0,1,0,0,2,1,0,1,1,1,2,0,2,2,1,0,2,1,1,1,2,1,2,1,2,2,2,1,1,2,2,0,2,0,0,2,2,1,2,2,1,2,0,2,2,1,2,1,2,1,0,0,0,2,2,2,2,1,2,2,2,0,1,2,2,2,2,2,2,2,2,2,2,0,0,2,2,2,1,2,1,2,0,0,2,0,0,2,2,0,0,2,0,1,1,1,2,1,0,2,2,2,0,2,2,2,2,0,0,2,1,2,0,1,1,2,2,1,0,2,2,0,2,2,1,2,1,1,2,0,2,0,0,0,2,2,0,2,2,2,2,2,2,2,2,2,0,2,2,2,0,2,0,1,2,2,2,0,0,1,2,0,2,2,2,0,2,0,0,0,2,0,1,2,2,2,2,2,0,1,1,2,1,0,2,0,1,2,2,2,1,2,0,1,2,2,1,2,2,0,2,2,0,2,2,2,2,2,2,2,2,0,2,2,0,2,2,2,2,1,1,1,2,1,1,2,2,2,1,2,1,0,1,0,2,2,0,2,2,0,1,0,1,2,2,1,2,1,1,2,2,2,2,1,2,2,2,1,1,2,2,0,2,0,2,2,2,1,1,2,2,0,2,2,2,2,2,2,2,1,2,0,2,2,2,2,0,1,2,2,2,2,2,2,2,2,0,1,2,0,1,0,1,1,2,1,0,1,1,2,2,1,2,0,1,0,2,0,0,1,2,2,0,2,2,2,2,2,2,1,1,0,1,1,1,0,1,0,2,2,0,1,1,0,2,2,0,2,2,2,1,0,0,2,2,2,2,0,1,2,2,2,2,2,2,2,1,2,2,2,2,2,2,2,0,2,2,2,1,2,2,1,2,0,2,1,1,2,2,1,1,2,2,2,2,2,2,1,2,0,2,2,0,1,2,2,2,1,2,2,2,2,2,1,1,2,2,1,1,0,2,2,2,2,0,2,1,0,2,2,1,2,2,0,2,2,2,0,1,2,2,2,2,1,2,2,1,0,2,0,1,2,1,2,2,0,1,1,1,1,2,2,2,0,1,2,0,2,1,2,2,1,1,1,2,2,2,1,2,1,0,2,1,0,2,0,2,1,2,2,1,1,0,2,2,2,1,0,2,2,0,2,2,2,0,2,2,2,0,2,2,0,1,2,0,2,0,2,2,2,0,2,2,0,0,2,2,2,0,2,1,1,0,0,2,1,1,2,2,0,2,2,0,2,1,1,2,1,2,1,2,2,1,1,2,2,2,0,0,2,0,1,1,0,0,2,2,2,0,0,2,0,2,2,2,2,2,2,2,0,1,1,0,1,2,2,0,2,1,2,1,2,2,2,2,2,1,2,2,1,2,2,1,0,2,1,1,2,0,0,2,0,2,2,2,1,2,1,2,0,2,0,1,0,2,2,2,1,2,0,2,2,2,0,1,2,1,2,2,2,2,1,0,0,2,2,2,1,2,1,1,0,0,0,0,2,2,0,2,2,2,2,0,2,1,0,1,2,2,1,2,1,2,2,1,2,2,2,1,2,1,0,0,0,2,1,1,2,2,1,1,0,0,1,0,1,2,1,2,2,2,2,2,0,2,2,2,2,0,2,2,2,2,0,2,2,2,2,1,2,2,1,0,2,2,2,1,0,2,0,0,1,1,0,2,2,0,1,2,2,1,2,2,2,2,2,1,2,2,0,1,2,2,1,2,0,2,0,2,0,2,0,2,2,2,2,1,0,1,2,2,1,1,2,2,2,2,1,0,0,1,2,0,0,1,1,2,2,2,2,2,2,2,0,0,0,2,2,2,0,1,0,2,1,2,2,2,2,0,2,2,0,2,2,2,0,0,2,2,0,1,2,2,0,0,2,2,2,0,2,0,0,0,2,2,0,2,2,2,1,2,2,1,0,2,2,2,0,1,1,1,2,2,1,2,2,2,0,2,2,0,2,1,1,0,2,2,2,1,2,2,2,1,2,1,2,2,2,1,2,2,1,2,1,2,2,2,2,1,2,0,2,2,2,2,2,1,2,2,1,0,1,2,2,2,1,2,2,2,2,2,2,2,2,2,1,2,1,0,2,2,0,1,1,2,2,2,1,0,0,1,2,0,1,1,1,2,2,2,2,2,1,2,1,1,1,2,2,2,2,0,2,1,1,2,2,2,1,1,2,2,1,2,2,1,0,2,2,2,0,2,1,2,1,1,2,2,2,0,2,1,0,0,0,2,1,2,2,2,2,0,2,0,1,1,0,2,2,0,2,1,2,2,0,2,2,2,0,0,2,1,0,1,0,0,2,2,0,2,1,2,2,1,2,0,2,1,0,2,0,2,2,1,0,2,2,0,2,2,2,0,2,2,0,2,2,0,0,2,2,2,1,2,1,2,1,0,1,1,2,2,2,2,1,0,2,0,0,0,1,2,2,0,1,2,2,2,1,1,0,0,2,1,0,0,1,2,2,2,2,2,2,2,1,0,0,2,0,2,0,0,2,1,1,0,2,2,2,1,2,2,0,0,2,0,2,0,2,2,1,2,0,2,0,0,2,2,2,1,2,2,0,0,1,2,0,2,2,2,2,2,2,0,0,1,0,2,0,1,2,0,2,0,1,2,2,2,1,1,2,1,0,1,2,0,2,2,2,2,2,0,1,2,2,1,2,1,2,0,2,1,0,0,0,2,2,2,2,1,2,1,2,2,2,2,2,2,0,2,1,2,1,1,0,2,1,1,2,1,2,2,0,2,0,0,1,0,0,0,2,2,2,0,0,2,2,2,2,1,0,1,1,0,2,2,1,1,2,2,2,0,0,0,2,0,0,1,1,1,0,0,2,0,2,0,2,2,1,1,2,2,1,2,2,0,0,1,2,2,2,2,0,2,2,0,2,2,2,2,2,1,2,2,2,2,2,2,2,2,2,2,2,0,0,1,2,2,0,0,0,2,2,1,1,2,2,2,1,2,2,0,2,2,0,1,2,2,1,2,0,2,2,2,1,2,2,2,1,1,1,1,0,0,2,2,2,0,0,0,2,1,2,2,1,2,2,0,1,2,2,0,1,2,1,2,2,2,2,2,0,2,0,2,0,1,2,2,2,2,1,2,0,1,1,2,2,2,0,2,0,1,0,0,2,0,1,2,0,2,2,0,2,2,0,0,0,0,1,0,1,2,2,2,1,1,2,2,1,1,2,2,1,2,2,2,2,2,2,2,2,0,0,2,1,0,2,2,2,0,2,1,0,2,1,1,0,1,2,2,1,2,0,1,2,1,0,2,2,2,0,2,2,2,2,2,2,2,2,0,2,1,2,1,0,0,2,2,0,0,1,0,0,1,2,2,2,2,0,1,1,1,1,0,1,2,2,1,1,2,2,1,2,2,1,2,2,2,1,2,0,2,0,0,0,2,2,0,2,2,1,2,0,2,0,1,2,0,0,2,2,2,1,2,2,2,2,2,2,2,0,1,2,2,0,2,2,2,0,2,2,1,2,0,0,2,2,0,1,2,2,0,1,1,1,0,2,2,1,0,0,2,2,1,2,1,2,2,2,2,2,1,1,2,0,1,2,2,2,1,2,1,2,0,1,0,1,2,2,2,0,0,2,0,0,0,0,2,1,1,1,2,2,0,1,2,2,2,0,1,2,2,1,0,2,2,2,2,1,1,2,2,2,1,2,2,2,1,1,0,0,1,0,0,0,2,2,1,1,1,2,1,2,2,1,2,2,2,0,2,0,2,2,2,2,2,0,0,2,0,0,2,0,2,2,1,2,1,2,1,1,2,1,1,2,2,2,2,1,1,0,1,2,2,1,0,1,0,0,2,2,2,0,1,0,0,1,0,0,1,1,0,2,1,2,0,2,2,2,1,0,2,0,0,2,0,2,0,0,2,2,2,2,2,1,0,2,2,2,0,2,1,1,2,0,2,2,0,2,2,2,1,2,1,1,1,2,2,0,0,2,0,2,2,1,2,0,2,0,1,2,0,0,1,2,0,2,2,2,1,2,0,0,2,1,0,2,0,2,1,1,2,1,1,2,2,2,1,0,0,2,0,2,2,1,2,0,2,2,2,2,2,2,1,0,2,0,1,0,1,2,2,2,2,1,2,1,1,0,1,2,2,1,2,1,2,2,2,1,2,1,0,2,1,0,1,2,2,0,2,2,1,1,1,1,1,2,2,0,0,1,2,2,2,1,0,2,2,0,0,0,2,0,0,2,2,1,2,2,2,0,0,2,0,1,0,2,2,2,2,2,2,1,2,2,0,1,2,2,2,2,1,2,0,0,1,1,2,2,1,1,0,0,1,1,2,2,2,0,2,2,1,2,2,2,0,2,2,0,0,1,1,1,2,0,0,2,2,0,2,0,0,1,1,0,2,2,2,2,2,2,2,2,2,1,2,2,2,0,2,0,1,2,0,2,2,2,1,0,1,1,2,0,2,0,2,0,0,2,0,0,1,0,1,0,2,2,2,2,2,1,1,1,0,2,1,2,1,0,2,2,1,1,1,2,2,1,0,1,0,1,1,2,0,1,2,2,2,2,1,0,2,0,2,2,1,0,0,2,2,1,0,0,2,2,0,2,2,2,2,2,1,0,0,2,1,2,1,2,2,0,0,0,0,1,1,0,2,2,2,0,0,2,2,1,2,2,2,2,0,2,1,1,0,0,2,2,2,0,1,1,2,1,2,0,2,2,0,0,1,1,2,2,2,1,2,2,1,1,1,2,0,2,2,2,2,2,2,1,2,1,2,2,1,2,2,1,1,1,1,1,2,0,2,2,1,0,1,0,0,1,0,1,2,1,2,2,2,2,2,1,2,1,2,0,1,2,0,2,2,2,1,0,2,1,1,2,0,1,0,0,0,0,0,0,2,2,2,1,0,2,2,1,0,2,0,2,2,1,2,0,0,2,2,0,2,2,2,2,1,2,2,1,0,0,2,1,0,2,2,2,1,2,1,2,2,1,2,2,1,1,2,0,1,1,2,1,2,1,0,2,1,0,0,2,1,2,2,0,1,1,0,1,1,0,1,2,0,2,1,1,1,1,1,0,2,2,2,2,2,2,0,2,0,2,0,1,0,1,2,0,1,2,1,2,2,1,0,1,1,0,2,0,2,1,1,0,0,0,0,2,1,2,1,0,2,2,2,2,2,1,1,2,1,1,0,1,2,0,2,2,2,1,0,1,0,2,2,2,2,0,1,0,2,2,0,2,2,1,1,0,2,0,0,2,2,0,1,2,2,2,0,0,0,2,1,2,2,2,1,2,1,2,0,0,2,1,1,2,2,2,0,1,1,1,0,1,2,2,1,1,2,0,1,1,2,0,2,2,1,2,0,0,1,0,0,2,2,0,1,2,1,2,0,0,0,2,0,1,0,2,2,0,1,0,2,2,1,0,1,2,0,2,0,0,1,2,1,2,2,0,2,0,2,1,2,1,1,2,1,1,2,2,2,0,1,0,0,0,1,2,2,1,2,1,2,2,2,1,0,1,1,2,0,1,1,2,1,0,2,2,2,0,2,2,0,0,2,2,1,2,0,2,0,2,2,2,2,1,2,1,2,1,0,2,2,2,1,2,2,0,0,2,2,1,2,2,2,2,0,2,1,1,2,2,2,2,2,2,2,0,0,2,1,0,1,2,2,0,2,2,1,2,1,0,2,2,2,1,2,2,1,2,1,1,1,2,2,2,2,2,0,1,2,0,0,2,2,2,1,0,2,0,0,1,2,2,1,2,0,1,1,2,1,0,2,1,2,0,2,1,0,2,2,2,2,0,0,0,2,0,2,1,2,2,1,1,0,2,2,1,0,1,2,0,2,2,2,2,1,1,0,2,1,1,1,0,1,2,2,2,0,2,2,2,2,0,1,1,2,2,1,1,0,0,0,2,2,0,2,1,2,1,2,2,0,1,0,2,2,1,2,1,1,2,0,2,2,2,2,2,0,1,2,0,0,2,1,2,2,0,0,0,2,0,0,1,2,2,0,1,0,2,1,1,2,2,2,2,0,1,0,1,2,0,1,2,2,1,0,2,1,0,1,0,1,2,1,2,1,2,2,1,2,1,2,2,1,1,1,1,2,2,0,0,1,0,0,0,2,2,0,2,2,0,2,1,1,0,2,1,2,1,2,2,2,0,0,1,1,1,0,0,0,2,2,2,2,0,2,0,1,2,0,1,0,0,2,1,2,2,2,2,1,2,1,2,1,0,0,2,1,2,2,2,0,2,2,1,1,1,2,0,1,2,0,2,0,1,2,1,1,0,1,2,1,2,2,2,0,2,1,1,1,2,2,2,2,2,2,2,0,0,2,0,0,1,2,0,0,1,1,0,0,1,2,0,0,2,2,2,1,1,2,2,1,2,2,1,2,2,2,2,2,1,0,2,0,0,0,1,2,2,2,0,2,2,2,0,2,0,2,2,1,2,2,0,2,1,2,2,1,0,0,1,2,2,0,2,2,1,2,2,2,2,2,2,2,0,1,0,2,0,0,0,2,2,2,0,2,1,2,0,1,0,0,0,1,1,2,2,2,0,2,2,0,0,1,0,1,0,1,0,0,2,2,2,2,0,2,1,2,1,1,2,2,1,0,2,2,1,1,2,2,0,1,2,2,2,1,2,2,1,0,1,0,1,0,2,2,2,2,0,2,2,1,2,2,1,1,1,1,0,1,0,2,2,2,0,0,1,0,2,2,0,0,2,2,2,1,2,0,0,0,1,0,2,2,2,1,2,0,2,2,1,2,2,1,1,0,1,2,2,1,0,1,0,1,1,2,1,2,2,1,1,2,0,0,1,0,0,1,1,2,0,0,2,2,2,2,1,2,0,0,2,2,2,2,1,0,0,1,1,2,2,0,0,0,2,2,2,0,1,0,1,1,1,1,2,2,2,2,2,1,1,1,2,2,0,1,0,2,0,0,2,0,1,2,1,2,2,0,1,1,2,1,2,2,2,0,2,1,2,1,2,0,2,2,2,2,2,1,2,0,2,2,2,2,0,2,1,0,0,1,1,2,0,0,1,1,0,1,0,2,2,1,2,2,1,0,1,0,2,0,1,0,2,0,1,0,2,0,0,1,0,2,2,1,1,2,1,2,2,0,0,0,2,2,0,2,2,0,0,1,0,2,1,2,1,1,1,0,1,2,2,1,0,1,2,0,0,2,2,0,2,2,2,2,2,0,0,1,0,2,0,2,0,2,1,2,2,1,0,1,1,1,1,0,0,1,2,0,0,2,1,1,2,2,1,2,0,2,0,0,2,1,1,0,0,2,2,0,1,0,0,0,2,2,2,1,2,1,1,1,1,1,0,1,2,2,1,1,2,1,1,2,0,2,1,1,1,2,2,1,1,2,1,1,1,0,0,0,0,1,2,2,2,2,2,1,0,0,2,1,1,0,2,0,0,0,0,0,2,1,2,2,2,0,0,2,0,0,1,1,2,1,2,0,2,2,2,0,2,2,0,2,1,1,2,2,1,0,1,2,0,0,2,1,2,1,2,1,0,0,0,2,2,2,2,1,0,1,2,2,1,2,0,1,2,2,2,0,1,0,0,2,0,2,0,0,1,2,1,1,1,0,2,2,0,0,2,2,2,0,2,2,0,2,2,2,1,0,0,1,2,0,2,2,2,0,2,1,0,1,1,1,2,2,2,1,2,2,2,0,2,0,1,0,0,2,1,1,2,2,0,0,0,2,0,2,1,1,0,2,1,2,0,0,1,2,1,0,2,0,2,0,0,1,1,1,1,1,1,0,1,0,0,0,0,2,1,1,0,1,0,1,0,2,2,1,1,2,2,1,1,1,2,0,2,0,2,0,2,0,2,0,2,1,1,2,0,1,2,2,2,1,1,0,1,1,1,0,1,1,1,1,1,0,1,0,0,2,2,1,0,1,0,1,2,2,2,2,1,2,0,2,1,0,0,2,2,0,1,0,0,2,2,1,2,0,1,0,1,2,1,0,0,1,0,1,0,1,0,1,0,0,1,2,0,}