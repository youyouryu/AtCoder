package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, k := io.NextInt(), io.NextInt()
	cnt := make([]int, n)
	for i := 0; i < k; i++ {
		di := io.NextInt()
		for j := 0; j < di; j++ {
			aij := io.NextInt() - 1
			cnt[aij]++
		}
	}
	ans := 0
	for i := range cnt {
		if cnt[i] == 0 {
			ans++
		}
	}
	io.Println(ans)
}
