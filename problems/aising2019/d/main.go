package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, q := io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	x := io.NextInts(q)
	ans := solve(a, x)
	io.Println(ans)
}

func solve(a, x []int) (ans int) {

	return
}
