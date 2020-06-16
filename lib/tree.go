package lib

// A Tree is a binary tree with integer values.
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

// NewTree creates new instance of Tree with value v
func NewTree(v int) *Tree {
	return &Tree{Value: v, Left: nil, Right: nil}
}

// Insert adds v to tree structure
func (t *Tree) Insert(v int) {
	if v < t.Value {
		if t.Left == nil {
			t.Left = NewTree(v)
		} else {
			t.Left.Insert(v)
		}
	} else {
		if t.Right == nil {
			t.Right = NewTree(v)
		} else {
			t.Right.Insert(v)
		}
	}
}

// Remove removes a value v from Tree
func (t *Tree) Remove(v int) *Tree {
	if v < t.Value {
		if t.Left != nil {
			t.Left = t.Left.Remove(v)
		}
		return t
	}
	if v > t.Value {
		if t.Right != nil {
			t.Right = t.Right.Remove(v)
		}
		return t
	}
	if t.Left != nil && t.Right != nil {
		successor := t.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		t.Right = t.Right.Remove(successor.Value)
		t.Value = successor.Value
		return t
	} else if t.Left != nil {
		return t.Left
	} else if t.Right != nil {
		return t.Right
	} else {
		return nil
	}
}
