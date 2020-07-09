package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	m, _ := sc.NextInt()
	h, _ := sc.NextInts(n)
	a, b := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		a[i], _ = sc.NextInt()
		b[i], _ = sc.NextInt()
	}
	ans := solve(n, m, h, a, b)
	fmt.Println(ans)
}

func solve(n, m int, h, a, b []int) (ans int) {
	lst := buildGraph(n, a, b)
	isGood := make([]bool, n)
	for i := range isGood {
		isGood[i] = true
	}
	for i := 0; i < n; i++ {
		if !isGood[i] {
			continue
		}
		for j := range lst[i] {
			if h[j] <= h[i] {
				isGood[j] = false
			} else {
				isGood[i] = false
			}
		}
	}
	return ans
}

func buildGraph(n int, a, b []int) [][]int {
	lst := make([][]int, n)
	for i := range a {
		lst[a[i]-1] = append(lst[a[i]-1], b[i]-1)
		lst[b[i]-1] = append(lst[b[i]-1], a[i]-1)
	}
	return lst
}
