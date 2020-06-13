package main

import (
	"fmt"
	"math"
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	m, _ := sc.NextInt()
	x, _ := sc.NextInt()
	c, a := []int{}, [][]int{}
	for i := 0; i < n; i++ {
		ci, _ := sc.NextInt()
		c = append(c, ci)
		ai, _ := sc.NextInts(m)
		a = append(a, ai)
	}
	ans := solve(n, m, x, c, a)
	fmt.Println(ans)
}

func solve(n, m, x int, c []int, a [][]int) (minCost int) {
	// i: combination
	// j: book
	// k: algorithm
	minCost = math.MaxInt64
	for i, maxIter := 0, 1<<n; i < maxIter; i++ {
		progress := make([]int, m)
		cost := 0
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				for k := 0; k < m; k++ {
					progress[k] += a[j][k]
				}
				cost += c[j]
			}
		}
		if !isOk(x, progress) {
			continue
		}
		if cost < minCost {
			minCost = cost
		}
	}
	if minCost == math.MaxInt64 {
		return -1
	}
	return minCost
}

func isOk(x int, progress []int) bool {
	for _, v := range progress {
		if v < x {
			return false
		}
	}
	return true
}
