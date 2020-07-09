package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	k, _ := sc.NextInt()
	a, _ := sc.NextInts(n)
	ans := solve(n, k, a)
	fmt.Println(ans)
}

func solve(n, k int, a []int) (ans int) {
	history := make([]int, n)
	for i := range history {
		history[i] = -1
	}
	current := 1
	history[current-1] = k
	for k > 0 {
		current = a[current-1]
		k--
		if history[current-1] != -1 {
			periodicTime := k - history[current-1]
			k %= periodicTime
		}
		history[current-1] = k
	}
	return current
}
