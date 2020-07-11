package main

import (
	"os"
	"sort"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	a, b, n := io.NextInt(), io.NextInt(), io.NextInt()
	max := b - 1
	if n < max {
		max = n
	}
	x := sort.Search(max, func(x int) bool {
		return f(a, b, x+1) < f(a, b, x)
	})
	io.Println(f(a, b, x))
}

func f(a, b, x int) int {
	return a*x/b - a*(x/b)
}
