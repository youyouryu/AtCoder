package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, k := io.NextInt(), io.NextInt()
	cnt := make([]int, k+1)
	for x := k; x >= 1; x-- {
		cnt[x] = lib.ModPow(k/x, n, mod)
		for y := 2 * x; y <= k; y += x {
			cnt[x] += mod - cnt[y]
			cnt[x] %= mod
		}
	}
	ans := 0
	for x := 1; x <= k; x++ {
		ans += x * cnt[x]
		ans %= mod
	}
	io.Println(ans)
}

const mod = 1e9 + 7
