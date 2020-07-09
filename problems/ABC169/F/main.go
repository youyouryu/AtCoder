package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	s, _ := sc.NextInt()
	a, _ := sc.NextInts(n)
	ans := solve(n, s, a)
	fmt.Println(ans)
}

const mod = 998244353

func solve(n, s int, a []int) (ans int) {
	dp, prev := make([]int, s+1), make([]int, s+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		copy(prev, dp)
		for j := 0; j <= s; j++ {
			dp[j] = 2 * prev[j]
			dp[j] %= mod
			if j-a[i] >= 0 {
				dp[j] += prev[j-a[i]]
				dp[j] %= mod
			}
		}
	}
	return dp[s]
}
