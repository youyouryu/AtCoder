package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = sc.NextInt()
		b[i], _ = sc.NextInt()
	}
	ans := solve(n, a, b)
	fmt.Println(ans)
}

func solve(n int, a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	if n%2 == 0 {
		m1 := b[n/2-1] - a[n/2-1] + 1
		m2 := b[n/2] - a[n/2] + 1
		return m1 + m2 - 1
	}
	return b[n/2] - a[n/2] + 1
}
