package collection

import "fmt"

// Tree represents a tree data structure with Push and Pop methods
type Tree struct {
	array []*Node
	index map[interface{}]int
}

// NewTree creates a new tree
func NewTree() *Tree {
	return &Tree{
		array: make([]*Node, 0),
		index: make(map[interface{}]int, 0),
	}
}

// Size returns the number of nodes in the tree
func (t *Tree) Size() int {
	return len(t.array)
}

// List gets the internal slice for the tree
func (t *Tree) List() []*Node {
	return t.array
}

// FindBase finds the first common base node for both nodes
func (t *Tree) FindBase(node1, node2 *Node) (base *Node) {
	n1path, n2path := node1.Path(), node2.Path()

	length := len(n1path)
	if len(n2path) < length {
		length = len(n2path)
	}

	for i := 0; i < length; i++ {
		if n1path[i] != n2path[i] {
			break
		}
		base = n1path[i]
	}

	return base
}

// FindByValue adds a new Node to the tree.
// It will also attach a reference to the parent
func (t *Tree) FindByValue(value interface{}) *Node {
	key, found := t.index[value]
	if !found {
		return nil
	}
	return t.array[key]
}

// Push adds a new Node to the tree
// It will not modify the node when adding.
func (t *Tree) Push(n *Node) *Node {
	arr := &t.array
	n.key = len(*arr)
	*arr = append(*arr, n)
	t.index[n.Value] = n.key
	return n
}

// Pop removes a Node from the Tree
func (t *Tree) Pop() *Node {
	old := t.array
	l := len(old)
	n := old[l-1]
	n.key = -1
	t.array = old[0:l-1]
	delete(t.index, n.Value)
	return n
}

// Print prints the tree to the console
func (t *Tree) Print() {
	for _, n := range t.array {
		fmt.Printf("<%v>", n)
	}
	fmt.Println()
}
