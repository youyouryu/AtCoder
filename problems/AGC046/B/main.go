package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	a, _ := sc.NextInt()
	b, _ := sc.NextInt()
	c, _ := sc.NextInt()
	d, _ := sc.NextInt()
	ans := solve(a, b, c, d)
	fmt.Println(ans)
}

const mod = 998244353

func solve(a, b, c, d int) (ans int) {
	dp := make([][]int, c+1)
	for i := range dp {
		dp[i] = make([]int, d+1)
	}
	dp[a][b] = 1
	for i := a; i <= c; i++ {
		for j := b; j <= d; j++ {
			if i == a && j == b {
				continue
			}
			tmp := j * dp[i-1][j] % mod
			tmp = (tmp + i*dp[i][j-1]%mod) % mod
			tmp = (tmp - (i-1)*(j-1)*dp[i-1][j-1]%mod + mod) % mod
			dp[i][j] = tmp
		}
	}
	return dp[c][d]
}
