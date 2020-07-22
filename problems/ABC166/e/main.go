package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	a := io.NextInts(n)
	cnt := map[int]int{}
	ans := 0
	for i := range a {
		ans += cnt[i-a[i]]
		cnt[i+a[i]]++
	}
	io.Println(ans)
}
