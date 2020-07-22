package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	a := io.NextInts(n - 1)
	cnt := map[int]int{}
	for i := range a {
		cnt[a[i]]++
	}
	for i := 1; i <= n; i++ {
		io.Println(cnt[i])
	}
}
