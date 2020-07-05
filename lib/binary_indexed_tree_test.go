package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryIndexedTree(t *testing.T) {
	n := 10
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	b := NewBinaryIndexedTree(n)
	for i := 0; i < n; i++ {
		b.Add(i+1, nums[i])
	}
	expected := make([]int, n)
	for i := 1; i < n; i++ {
		expected[i] = expected[i-1] + nums[i]
	}
	for i := 0; i < n; i++ {
		assert.Equal(t, expected[i], b.Sum(i+1))
	}
}
