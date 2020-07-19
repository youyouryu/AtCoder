package lib

import "fmt"

// Matrix is a 2d matrix of integers
type Matrix [][]int

// NewMatrix returns n*m matrix filled with 0
func NewMatrix(n, m int) Matrix {
	ret := make(Matrix, n)
	for i := range ret {
		ret[i] = make([]int, m)
	}
	return ret
}

// NewIdentityMatrix returns a n*n matrix whose diagonal elements are 0
func NewIdentityMatrix(n int) Matrix {
	ret := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				ret[i][j] = 1
			}
		}
	}
	return ret
}

// Mul calculates matrix product
func (a Matrix) Mul(b Matrix) Matrix {
	if len(a[0]) != len(b) {
		panic(fmt.Errorf("shape mismatch"))
	}
	n, m, l := len(a), len(b), len(b[0])
	ret := NewMatrix(n, l)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < l; k++ {
				ret[i][k] += a[i][j] * b[j][k]
			}
		}
	}
	return ret
}

// Pow calculates matrix power
func (a Matrix) Pow(k int) Matrix {
	n := len(a)
	if len(a) != len(a[0]) {
		panic(fmt.Errorf("a must be square Matrix"))
	}
	ret := NewIdentityMatrix(n)
	for k > 0 {
		if k&1 == 1 {
			ret = ret.Mul(a)
		}
		a = a.Mul(a)
		k >>= 1
	}
	return ret
}
