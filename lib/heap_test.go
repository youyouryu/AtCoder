package lib

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	q := &PriorityQueue{}
	pairs := []Pair{{2, 2}, {1, 3}, {1, 2}, {2, 3}}
	for i := range pairs {
		heap.Push(q, pairs[i])
	}
	expected := []Pair{{1, 2}, {1, 3}, {2, 2}, {2, 3}}
	for i := range expected {
		assert.Equal(t, expected[i], heap.Pop(q).(Pair))
	}
}

type Pair struct {
	a, b int
}

func (p Pair) Less(q Comparable) bool {
	if p.a == q.(Pair).a {
		return p.b < q.(Pair).b
	}
	return p.a < q.(Pair).a
}
