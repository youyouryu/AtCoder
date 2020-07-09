package lib

// MergeFunc merges two.Nodes to one.
type MergeFunc func(interface{}, interface{}) interface{}

// SegmentTree is data structure which can calculate the representative value of section
type SegmentTree struct {
	Offset    int
	Nodes     []interface{}
	MergeFunc MergeFunc
}

// NewSegmentTree returns a new initialized instance of SegmentTree
func NewSegmentTree(data []interface{}, mergeFunc MergeFunc) *SegmentTree {
	depth := 0
	for 1<<depth < len(data) {
		depth++
	}
	size := 1<<(depth+1) - 1
	Nodes := make([]interface{}, size)
	offset := 1<<depth - 1
	for i := offset; i < offset+len(data); i++ {
		Nodes[i] = data[i-offset]
	}
	s := &SegmentTree{
		Offset:    offset,
		Nodes:     Nodes,
		MergeFunc: mergeFunc,
	}
	return s.fix()
}

func (s *SegmentTree) fix() *SegmentTree {
	for i := s.Offset - 1; i >= 0; i-- {
		left, right := s.Nodes[2*i+1], s.Nodes[2*i+2]
		if right == nil {
			s.Nodes[i] = left
			continue
		}
		s.Nodes[i] = s.MergeFunc(left, right)
	}
	return s
}

// Query returns a merged node of [begin, end)
func (s *SegmentTree) Query(begin, end int) interface{} {
	return s.query(begin, end, 0, 0, len(s.Nodes)-s.Offset)
}

func (s *SegmentTree) query(begin, end, k, cBegin, cEnd int) interface{} {
	if cEnd <= begin || end <= cBegin {
		return nil
	} else if begin <= cBegin && cEnd <= end {
		return s.Nodes[k]
	}
	c1 := s.query(begin, end, 2*k+1, cBegin, (cBegin+cEnd)/2)
	c2 := s.query(begin, end, 2*k+2, (cBegin+cEnd)/2, cEnd)
	if c1 == nil {
		return c2
	} else if c2 == nil {
		return c1
	}
	return s.MergeFunc(c1, c2)
}

// Update updates a leaf node and reconstruct tree structure.
func (s *SegmentTree) Update(index int, value interface{}) *SegmentTree {
	s.Nodes[index] = value
	return s.fix()
}
