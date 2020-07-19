package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMulMatrix(t *testing.T) {
	a := Matrix([][]int{{0, 1, 2}, {3, 4, 5}})
	b := Matrix([][]int{{0, 1}, {2, 3}, {4, 5}})
	c := a.Mul(b)
	expected := Matrix([][]int{{10, 13}, {28, 40}})
	for i := range expected {
		for j := range expected[i] {
			assert.Equal(t, expected[i][j], c[i][j])
		}
	}
}

func TestPowMatrix(t *testing.T) {
	a := Matrix([][]int{{1, 1}, {0, 1}})
	b := a.Pow(3)
	expected := Matrix([][]int{{1, 3}, {0, 1}})
	for i := range expected {
		for j := range expected[i] {
			assert.Equal(t, expected[i][j], b[i][j])
		}
	}
}
