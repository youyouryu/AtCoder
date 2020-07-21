package lib

// IntHeap is heap for interger
// gollect: keep methods
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to heap
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop returns the minimum element of the heap and remove it
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Comparable interface{ Less(Comparable) bool }

// PriorityQueue
// gollect: keep methods
type PriorityQueue []Comparable

func (q PriorityQueue) Len() int           { return len(q) }
func (q PriorityQueue) Less(i, j int) bool { return q[i].Less(q[j]) }
func (q PriorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

// Push adds an element
func (q *PriorityQueue) Push(x interface{}) { *q = append(*q, x.(Comparable)) }

// Pop gets and remove the minimum element
func (q *PriorityQueue) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}
