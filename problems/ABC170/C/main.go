package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	x, _ := sc.NextInt()
	n, _ := sc.NextInt()
	p, _ := sc.NextInts(n)
	ans := solve(x, n, p)
	fmt.Println(ans)
}

func solve(x, n int, p []int) int {
	sort.Slice(p, func(i, j int) bool { return p[i] < p[j] })
	for i := 0; ; i++ {
		if !contains(p, x-i) {
			return x - i
		} else if !contains(p, x+i) {
			return x + i
		}
	}
}

func contains(p []int, v int) bool {
	n := len(p)
	i := sort.Search(n, func(i int) bool { return p[i] >= v })
	return i < n && p[i] == v
}
