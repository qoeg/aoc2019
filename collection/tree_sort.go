package collection

import "fmt"

// ByValueString is a sorting type wrapper for Tree.
// It will sort the Tree by the string representation of the node values.
type ByValueString Tree

func (t ByValueString) Len() int {
	return len(t.array)
}

func (t ByValueString) Swap(i, j int) {
	arr, index := t.array, t.index
	arr[i], arr[j] = arr[j], arr[i]
	arr[i].key = i
	index[arr[i].Value] = i
	arr[j].key = j
	index[arr[j].Value] = j
}

func (t ByValueString) Less(i, j int) bool {
	iStr, jStr := fmt.Sprint(t.array[i].Value), fmt.Sprint(t.array[j].Value)
	if iStr == jStr {
		return iStr < jStr
	}
	return iStr < jStr
}

// Push pushes a new Node onto the underlying Tree
func (t *ByValueString) Push(x interface{}) {
	(*Tree)(t).Push(x.(*Node))
}

// Pop pops a Node off the underlying Tree and returns it
func (t *ByValueString) Pop() interface{} {
	return (*Tree)(t).Pop()
}