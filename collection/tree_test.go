package collection

import (
	"reflect"
	"testing"
)

func TestTree_Push(t *testing.T) {
	tree := NewTree()
	node := NewNode("test", nil)

	type args struct {
		n *Node
	}
	tests := []struct {
		name string
		t    *Tree
		args args
		want *Node
	}{
		{"Test 1", tree, args{node}, node},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Push(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_FindByValue(t *testing.T) {
	tree := NewTree()
	nodeA := NewNode("A", nil)
	nodeB := NewNode("B", nil)
	tree.Push(nodeA)
	tree.Push(nodeB)

	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		t    *Tree
		args args
		want *Node
	}{
		{"Test 1", tree, args{"A"}, nodeA},
		{"Test 2", tree, args{"B"}, nodeB},
		{"Test 3", tree, args{"C"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.FindByValue(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.FindByValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
