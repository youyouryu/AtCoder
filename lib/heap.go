package lib

// Pair has id and value
type Pair struct {
	Key   int
	Value int
}

// PriorityQueue is heap for Pair
// gollect: keep methods
type PriorityQueue []Pair

// Len is used int heap interface
func (q PriorityQueue) Len() int { return len(q) }

// Less is used int heap interface
func (q PriorityQueue) Less(i, j int) bool { return q[i].Value < q[j].Value }

// Swap is used int heap interface
func (q PriorityQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

// Push adds an element to heap
func (q *PriorityQueue) Push(x interface{}) { *q = append(*q, x.(Pair)) }

// Pop returns the minimum element from heap
func (q *PriorityQueue) Pop() interface{} {
	old := *q
	*q = old[:len(old)-1]
	return old[len(old)-1]
}
