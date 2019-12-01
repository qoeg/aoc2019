package grid

// Tree represents a tree data structure with Push and Pop methods
type Tree []*Node

// Push adds a new Node to the tree
func (t *Tree) Push(n *Node) {
	n.index = len(*t)
	*t = append(*t, n)
}

// Pop removes a Node from the Tree
func (t *Tree) Pop() *Node {
	old := *t
	l := len(old)
	n := old[l-1]
	n.index = -1
	*t = old[0:l-1]
	return n
}

// NewTree creates a new Tree
func NewTree(root Cell) *Tree {
	t := make(Tree, 1)
	t[0] = &Node{
		parent: nil,
		children: []*Node{},
		Value: root,
	}
	return &t
}
