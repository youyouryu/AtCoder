package main

import (
	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo()
	defer io.Flush()
	k := io.NextInt()
	s := io.Next()
	ans := solve(k, s)
	io.Println(ans)
}

const mod = 1e9 + 7
const nChars = 26

func solve(k int, s string) (ans int) {
	n := len(s)
	cmb := lib.NewModCombination(mod)
	cmb.Calc(n-1+k, n-1+k)
	ans = 0
	for i := 0; i <= k; i++ {
		now := cmb.Calc(n-1+i, n-1)
		now = lib.ModMul(now, lib.ModPow(25, i, mod), mod)
		now = lib.ModMul(now, lib.ModPow(26, k-i, mod), mod)
		ans = (ans + now) % mod
	}
	return
}
