package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, m := io.NextInt(), io.NextInt()
	a := io.NextInts(m)
	for i := range a {
		n -= a[i]
	}
	if n < 0 {
		io.Println(-1)
	} else {
		io.Println(n)
	}
}
