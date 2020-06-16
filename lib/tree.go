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
func (t *Tree) Remove(v int) {
	var parent, current *Tree
	parent, current = nil, t
	for {
		if current.Value > v {
			if current.Left == nil {
				return
			}
			parent, current = current, current.Left
		} else if current.Value < v {
			if current.Right == nil {
				return
			}
			parent, current = current, current.Right
		} else {
			if current.Left != nil && current.Right != nil {
				afterChild := current.Right
				for afterChild.Left != nil {
					afterChild = afterChild.Left
				}
				current.Remove(afterChild.Value)
				current.Value = afterChild.Value
			} else if current.Left != nil {
				connect(parent, current, current.Left)
			} else if current.Right != nil {
				connect(parent, current, current.Right)
			} else {
				connect(parent, current, nil)
			}
			return
		}
	}
}

func connect(parent, beforeChild, afterChild *Tree) {
	if parent.Left == beforeChild {
		parent.Left = afterChild
	} else {
		parent.Right = afterChild
	}
}
