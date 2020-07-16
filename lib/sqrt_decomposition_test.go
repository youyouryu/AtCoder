package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrtDecomposition(t *testing.T) {
	n := 10
	s := NewSqrtDecomposition(n)
	for i := 0; i < n; i++ {
		s.Add(i, i)
	}
	begin, end := []int{0, 0, 0}, []int{4, 5, 10}
	expected := []int{6, 10, 45}
	for i := range expected {
		assert.Equal(t, expected[i], s.GetSum(begin[i], end[i]))
	}
}
