package grid

// TreeByReadOrder is a sorting type wrapper for Tree.
// It will sort the Tree by the Cell value of each Node,
// in coordinate "read order", left-to-right, top-to-bottom.
type TreeByReadOrder Tree

func (t TreeByReadOrder) Len() int {
	return len(t)
}

func (t TreeByReadOrder) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
	t[i].index = i
	t[j].index = j
}

func (t TreeByReadOrder) Less(i, j int) bool {
	if t[i].Value.Pos().Y == t[j].Value.Pos().Y {
		return t[i].Value.Pos().X < t[j].Value.Pos().X
	}
	return t[i].Value.Pos().Y < t[j].Value.Pos().Y
}

// Push pushes a new Node onto the underlying Tree
func (t *TreeByReadOrder) Push(x interface{}) {
	(*Tree)(t).Push(x.(*Node))
}

// Pop pops a Node off the underlying Tree and returns it
func (t *TreeByReadOrder) Pop() interface{} {
	return (*Tree)(t).Pop()
}