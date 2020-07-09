package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	a, b := make([]int, n-1), make([]int, n-1)
	for i := 0; i < n-1; i++ {
		ai, bi := io.NextInt(), io.NextInt()
		if ai > bi {
			ai, bi = bi, ai
		}
		a[i], b[i] = ai, bi
	}
	ans := solve(n, a, b)
	io.Println(ans)
}

// O(n)
// 連結成分数 = 頂点数 - 辺数 (閉路がないため)
func solve(n int, a, b []int) (ans int) {
	v, e := 0, 0
	for i := 1; i <= n; i++ {
		v += i * (n - i + 1)
	}
	for i := 0; i < n-1; i++ {
		e += a[i] * (n - b[i] + 1)
	}
	return v - e
}
