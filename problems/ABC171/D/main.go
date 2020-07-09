package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	a, _ := sc.NextInts(n)
	q, _ := sc.NextInt()
	b, c := make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		b[i], _ = sc.NextInt()
		c[i], _ = sc.NextInt()
	}
	ans := solve(n, q, a, b, c)
	for _, v := range ans {
		fmt.Println(v)
	}
}

func solve(n, q int, a, b, c []int) (ans []int) {
	count := map[int]int{}
	for _, v := range a {
		count[v]++
	}
	ans = make([]int, q+1)
	ans[0] = sum(count)
	for i := 0; i < q; i++ {
		ans[i+1] = ans[i] - b[i]*count[b[i]] + c[i]*count[b[i]]
		count[c[i]] += count[b[i]]
		delete(count, b[i])
	}
	return ans[1:]
}

func sum(count map[int]int) int {
	ret := 0
	for k, v := range count {
		ret += k * v
	}
	return ret
}
