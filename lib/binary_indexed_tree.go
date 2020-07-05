package lib

// BinaryIndexedTree calculates sum of specific section
type BinaryIndexedTree []int

// NewBinaryIndexedTree returns new instance of BinaryIndexedTree
func NewBinaryIndexedTree(size int) *BinaryIndexedTree {
	b := make(BinaryIndexedTree, size+1)
	return &b
}

// Sum retuns sum of [0, i)
func (b *BinaryIndexedTree) Sum(i int) int {
	sum := 0
	for i > 0 {
		sum += (*b)[i]
		i -= i & -i
	}
	return sum
}

// Add adds x to i th value
func (b *BinaryIndexedTree) Add(i, x int) *BinaryIndexedTree {
	for i < len(*b) {
		(*b)[i] += x
		i += i & -i
	}
	return b
}
