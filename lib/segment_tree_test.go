package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuerySegmentTree(t *testing.T) {
	n := 6
	data := prepareData(n)
	st := NewSegmentTree(data, merge)
	begin := []int{0, 0, 0, 0, 0, 1}
	end := []int{1, 2, 3, 4, 5, 8}
	for i := range begin {
		expected := calcExpected(data, begin[i], end[i])
		out := st.Query(begin[i], end[i])
		assert.Equal(t, expected, out)
	}
}

func TestUpdateSegmentTree(t *testing.T) {
	n := 6
	data := prepareData(n)
	st := NewSegmentTree(data, merge)
	st.Update(st.Offset+2, 102)
	assert.Equal(t, 115, st.Nodes[0])
}

func prepareData(n int) []interface{} {
	data := make([]interface{}, n)
	for i := range data {
		data[i] = i
	}
	return data
}

func merge(a, b interface{}) interface{} {
	na, nb := a.(int), b.(int)
	return na + nb
}

func calcExpected(data []interface{}, begin, end int) int {
	sum := 0
	for i := begin; i < end; i++ {
		if i >= len(data) {
			break
		}
		sum += data[i].(int)
	}
	return sum
}
