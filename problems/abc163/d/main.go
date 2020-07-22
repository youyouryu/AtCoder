package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

const mod = 1e9 + 7

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, k := io.NextInt(), io.NextInt()
	cnt := 0
	for i := k; i <= n+1; i++ {
		cnt += (i*(n-i+1) + 1) % mod
		cnt %= mod
	}
	io.Println(cnt)
}
