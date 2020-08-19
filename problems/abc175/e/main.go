package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	R, C, K := io.NextInt(), io.NextInt(), io.NextInt()
	values := make([][]int, R)
	for i := range values {
		values[i] = make([]int, C)
	}
	for i := 0; i < K; i++ {
		ri, ci, vi := io.NextInt()-1, io.NextInt()-1, io.NextInt()
		values[ri][ci] = vi
	}
	ans := solve(values)
	io.Println(ans)
}

func solve(values [][]int) int {
	R, C, L := len(values), len(values[0]), 3
	dp := make([][][]int, R)
	for i := range dp {
		dp[i] = make([][]int, C)
		for j := range dp[i] {
			dp[i][j] = make([]int, L+1)
		}
	}
	dp[0][0][1] = values[0][0]
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			for l := 0; l < L+1; l++ {
				if i+1 < R {
					dp[i+1][j][0] = lib.Max(dp[i+1][j][0], dp[i][j][l])
					dp[i+1][j][1] = lib.Max(dp[i+1][j][1], dp[i][j][l]+values[i+1][j])
				}
				if j+1 < C {
					dp[i][j+1][l] = lib.Max(dp[i][j+1][l], dp[i][j][l])
					if l+1 < L+1 {
						dp[i][j+1][l+1] = lib.Max(dp[i][j+1][l+1], dp[i][j][l]+values[i][j+1])
					}
				}
			}
		}
	}
	return lib.Max(dp[R-1][C-1]...)
}
