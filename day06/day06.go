package day06

import (
	"strconv"
	"strings"
	"github.com/qoeg/aoc2019/collection"
)

var tree *collection.Tree

func init() {
	tree = collection.NewTree()
	tree.Push(collection.NewNode("COM", nil))

	fill(tree, Input)
}

func fill(tree *collection.Tree, input []string) {
	scan := make([]string, len(input))
	copy(scan, input)

	for len(scan) > 0 {
		nextScan := []string{}

		for _, orbit := range scan {
			bodies := strings.Split(orbit,")")
			parent := tree.FindByValue(bodies[0])
			if parent == nil {
				nextScan = append(nextScan, orbit)
				continue
			}
			node := collection.NewNode(bodies[1], parent)
			tree.Push(node)
		}

		scan = make([]string, len(nextScan))
		copy(scan, nextScan)
	}
}

func countMany(nodes []*collection.Node) (count int) {
	for _, n := range nodes {
		count += countOrbits(n)
	}
	return count
}

func countOrbits(node *collection.Node) (count int) {
	p := node.Parent()
	for p != nil {
		count++
		p = p.Parent()
	}
	return count
}

// Answer1 returns the first puzzle answer
func Answer1() string {
	count := countMany(tree.List())
	return strconv.Itoa(count)
}

// Answer2 returns the second puzzle answer
func Answer2() string {
	you := tree.FindByValue("YOU")
	san := tree.FindByValue("SAN")
	base := tree.FindBase(you, san)

	count := countOrbits(you.Parent())
	count += countOrbits(san.Parent())
	count -= 2*countOrbits(base)

	return strconv.Itoa(count)
}
