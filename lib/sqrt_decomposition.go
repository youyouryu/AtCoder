package lib

import "math"

// SqrtDecomposition can calculate sum of [begin, end) with O(√n)
type SqrtDecomposition struct {
	Data   []int
	Bucket []int
}

// NewSqrtDecomposition returns a new instannce of SqrtDecomposition
func NewSqrtDecomposition(n int) *SqrtDecomposition {
	sqrtN := int(math.Sqrt(float64(n)))
	data := make([]int, n)
	bucket := make([]int, n/sqrtN+1)
	return &SqrtDecomposition{
		Data:   data,
		Bucket: bucket,
	}
}

// Add adds x to ith value
func (s *SqrtDecomposition) Add(i, x int) {
	s.Data[i] += x
	sqrtN := int(math.Sqrt(float64(len(s.Data))))
	s.Bucket[i/sqrtN] += x
}

// GetSum returns sum of [begin, end)
func (s *SqrtDecomposition) GetSum(begin, end int) (sum int) {
	n := len(s.Data)
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 0; i < n/sqrtN+1; i++ {
		l, r := i*sqrtN, (i+1)*sqrtN
		if r <= begin || end <= l {
			continue
		}
		if begin <= l && r <= end {
			sum += s.Bucket[i]
		} else {
			for j := Max(begin, l); j < Min(end, r); j++ {
				sum += s.Data[j]
			}
		}
	}
	return
}
