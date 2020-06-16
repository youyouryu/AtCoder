package lib

import (
	"testing"
)

func TestInsert(t *testing.T) {
	ref := []int{10, 5, 3, 7, 6, 8, 9, 15}
	tree := initTree(ref)
	values := *dfsOrder(tree, nil)
	assertSlieceEqual(values, ref, t)
}

func initTree(values []int) *Tree {
	tree := NewTree(values[0])
	for _, v := range values[1:] {
		tree.Insert(v)
	}
	return tree
}

func dfsOrder(tree *Tree, values *[]int) *[]int {
	if values == nil {
		values = &[]int{}
	}
	*values = append(*values, tree.Value)
	if tree.Left != nil {
		dfsOrder(tree.Left, values)
	}
	if tree.Right != nil {
		dfsOrder(tree.Right, values)
	}
	return values
}

func assertSlieceEqual(a, b []int, t *testing.T) {
	if len(a) != len(b) {
		t.Fatal("length mismatch")
	}
	for i := range a {
		if a[i] != b[i] {
			t.Fatal("value mismatch")
		}
	}
}

// Remove node which has no childeren
func TestRemove1(t *testing.T) {
	tree := initTree([]int{10, 5, 3, 7, 6, 8, 9, 15})
	tree.Remove(15)
	ref := []int{10, 5, 3, 7, 6, 8, 9}
	values := *dfsOrder(tree, nil)
	assertSlieceEqual(values, ref, t)
}

// Remove node which has 1 child
func TestRemove2(t *testing.T) {
	tree := initTree([]int{10, 5, 3, 7, 6, 8, 9, 15})
	tree = tree.Remove(8)
	ref := []int{10, 5, 3, 7, 6, 9, 15}
	values := *dfsOrder(tree, nil)
	assertSlieceEqual(values, ref, t)
}

// Remove node which has 2 childeren
func TestRemove3(t *testing.T) {
	tree := initTree([]int{10, 5, 3, 7, 6, 8, 9, 15})
	tree = tree.Remove(7)
	ref := []int{10, 5, 3, 8, 6, 9, 15}
	values := *dfsOrder(tree, nil)
	assertSlieceEqual(values, ref, t)
}

// Remove node which does not exist
func TestRemove4(t *testing.T) {
	tree := initTree([]int{10, 5, 3, 7, 6, 8, 9, 15})
	tree = tree.Remove(99)
	ref := []int{10, 5, 3, 7, 6, 8, 9, 15}
	values := *dfsOrder(tree, nil)
	assertSlieceEqual(values, ref, t)
}
