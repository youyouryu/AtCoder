package lib

import "fmt"

// MergeFunc merges two nodes to one.
type MergeFunc func(interface{}, interface{}) interface{}

// SegmentTree is data structure which can calculate the representative value of section
type SegmentTree struct {
	offset    int
	nodes     []interface{}
	mergeFunc MergeFunc
}

// NewSegmentTree returns a new initialized instance of SegmentTree
func NewSegmentTree(data []interface{}, mergeFunc MergeFunc) *SegmentTree {
	depth := 0
	for 1<<depth < len(data) {
		depth++
	}
	size := 1<<(depth+1) - 1
	nodes := make([]interface{}, size)
	offset := 1<<depth - 1
	for i := offset; i < offset+len(data); i++ {
		nodes[i] = data[i-offset]
	}
	s := &SegmentTree{
		offset:    offset,
		nodes:     nodes,
		mergeFunc: mergeFunc,
	}
	return s.fix()
}

func (s *SegmentTree) fix() *SegmentTree {
	for i := s.offset - 1; i >= 0; i-- {
		left, right := s.nodes[2*i+1], s.nodes[2*i+2]
		if right == nil {
			s.nodes[i] = left
			continue
		}
		s.nodes[i] = s.mergeFunc(left, right)
	}
	return s
}

// Calc returns a merged node of [begin, end)
func (s *SegmentTree) Calc(begin, end int) interface{} {
	return s.calc(begin, end, 0, 0, len(s.nodes)-s.offset)
}

func (s *SegmentTree) calc(begin, end, k, cBegin, cEnd int) interface{} {
	if cEnd <= begin || end <= cBegin {
		return nil
	} else if begin <= cBegin && cEnd <= end {
		return s.nodes[k]
	}
	c1 := s.calc(begin, end, 2*k+1, cBegin, (cBegin+cEnd)/2)
	c2 := s.calc(begin, end, 2*k+2, (cBegin+cEnd)/2, cEnd)
	if c1 == nil {
		return c2
	} else if c2 == nil {
		return c1
	}
	return s.mergeFunc(c1, c2)
}

// Get returns a node
func (s *SegmentTree) Get(index int) interface{} {
	return s.nodes[index]
}

// Top returns the root node
func (s *SegmentTree) Top() interface{} {
	return s.Get(0)
}

// Update updates a leaf node and reconstruct tree structure.
func (s *SegmentTree) Update(index int, value interface{}) *SegmentTree {
	s.nodes[index] = value
	return s.fix()
}

// Print outputs contens of the tree for debug
func (s *SegmentTree) Print() {
	fmt.Printf("offset=%d\n", s.offset)
	for i, v := range s.nodes {
		fmt.Printf("%d %+v\n", i, v)
	}
	fmt.Println()
}
