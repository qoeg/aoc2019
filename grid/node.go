package grid

import "fmt"

// Node represents a node in a tree.
// It has one parent and mulitple children.
// The Value of a Node is a Grid Cell.
type Node struct{
	index int
	parent *Node
	children []*Node
	Value Cell
}

// Path returns the slice of Cells from the root to the Node
func (n *Node) Path() []Cell {
	path := []Cell{n.Value}

	p := n.parent
	for p != nil {
		path = append([]Cell{p.Value}, path...)
		p = p.parent
	}

	return path
}

// Pos gets the underlying Cell value's coordinate position
func (n *Node) Pos() Coordinate {
	return n.Value.Pos()
}

// String gets a string representing the Node
func (n *Node) String() string {
	p := "root"
	if n.parent != nil {
		p = "node"
	}
	return fmt.Sprintf("<P[%s] {%d,%d}>", p, n.Value.Pos().X, n.Value.Pos().Y)
}

// NewNode creates a new Node
func NewNode(c Cell, p *Node) *Node {
	n := &Node{
		parent: p,
		children: []*Node{},
		Value: c,
	}
	p.children = append(p.children, n)
	return n
}
