package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	s, _ := sc.NextString()
	k, _ := sc.NextInt()
	ans := solve(s, k)
	fmt.Println(ans)
}

const mod = 998244353

func solve(s string, m int) (ans int) {
	count := countOnes(s)
	n := len(count)
	sum := 0
	for _, v := range count {
		sum += v
	}
	m = lib.Min(m, sum)
	dp := initDp(m)
	dp[0][0] = 1
	for i := 0; i < n; i++ {
		tmp := initDp(m)
		for j := 0; j <= m; j++ {
			for k := 0; k <= m; k++ {
				// do nothing
				tmp[j][k] += dp[j][k]
				tmp[j][k] %= mod
				// move from stock
				for l := 1; l <= k; l++ {
					tmp[j][k-l] += dp[j][k]
					tmp[j][k-l] %= mod
				}
				// move to stock
				for l := 1; l <= count[n-1-i] && j+l <= m && k+l <= m; l++ {
					tmp[j+l][k+l] += dp[j][k]
					tmp[j+l][k+l] %= mod
				}
			}
		}
		copy(dp, tmp)
	}
	for j := range dp {
		ans += dp[j][0]
		ans %= mod
	}
	return ans
}

func countOnes(s string) []int {
	ret := []int{}
	count := 0
	for _, c := range s {
		if c == '1' {
			count++
			continue
		}
		ret = append(ret, count)
		count = 0
	}
	ret = append(ret, count)
	return ret
}

func initDp(m int) [][]int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	return dp
}
