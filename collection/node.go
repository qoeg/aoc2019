package collection

import "fmt"

// Node represents a node in a tree.
// It holds a value of type interface{}.
// It has one parent and mulitple children.
type Node struct{
	key int
	parent *Node
	children []*Node
	Value interface{}
}

// NewNode creates a new Node
func NewNode(value interface{}, parent *Node) *Node {
	n := &Node{
		parent: parent,
		children: []*Node{},
		Value: value,
	}

	if parent != nil {
		parent.children = append(parent.children, n)
	}
	
	return n
}

// Parent returns the parent of the node
func (n *Node) Parent() *Node {
	return n.parent
}

// Path returns the slice of Nodes from the given base to the node.
// A base node of nil will get the path to the root.
func (n *Node) Path() []*Node {
	path := []*Node{n}

	p := n.parent
	for p != nil {
		path = append([]*Node{p}, path...)
		p = p.parent
	}

	return path
}

// String gets a string representing the Node
func (n *Node) String() string {
	p := "root"
	if n.parent != nil {
		p = fmt.Sprint(n.parent.Value)
	}
	return fmt.Sprintf("%v <- %v", p, n.Value)
}
