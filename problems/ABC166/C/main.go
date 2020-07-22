package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, m := io.NextInt(), io.NextInt()
	h := io.NextInts(n)
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		from, to := io.NextInt()-1, io.NextInt()-1
		g[from] = append(g[from], to)
		g[to] = append(g[to], from)
	}
	ans := solve(h, g)
	io.Println(ans)
}

func solve(h []int, g [][]int) (ans int) {
	for i := range g {
		isGood := true
		for j := range g[i] {
			from, to := i, g[i][j]
			if h[from] <= h[to] {
				isGood = false
				break
			}
		}
		if isGood {
			ans++
		}
	}
	return ans
}
