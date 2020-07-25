package main

import (
	"math"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, x, y := io.NextInt(), io.NextInt()-1, io.NextInt()-1
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ij := j - i
			ixyj := lib.Abs(x-i) + 1 + lib.Abs(y-j)
			iyxj := lib.Abs(y-i) + 1 + lib.Abs(x-j)
			d := lib.Min(ij, ixyj, iyxj)
			dist[d]++
		}
	}
	for i := 1; i < n; i++ {
		io.Println(dist[i])
	}
}

// O(n^3): TLE
func main2() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, x, y := io.NextInt(), io.NextInt()-1, io.NextInt()-1
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i == j {
				continue
			}
			dp[i][j] = inf
		}
	}
	for i := 0; i < n; i++ {
		if i-1 >= 0 {
			dp[i][i-1] = 1
		}
		if i+1 < n {
			dp[i][i+1] = 1
		}
	}
	dp[x][y], dp[y][x] = 1, 1

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dp[i][k] == inf || dp[k][j] == inf {
					continue
				}
				dp[i][j] = lib.Min(dp[i][j], dp[i][k]+dp[k][j])
			}
		}
	}
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cnt[dp[i][j]]++
		}
	}
	for i := 1; i < n; i++ {
		io.Println(cnt[i])
	}
}

const inf = math.MaxInt64
