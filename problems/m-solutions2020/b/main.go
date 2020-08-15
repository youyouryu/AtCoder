package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	a, b, c := io.NextInt(), io.NextInt(), io.NextInt()
	k := io.NextInt()
	cnt := 0
	for a >= b {
		b *= 2
		cnt++
	}
	for b >= c {
		c *= 2
		cnt++
	}
	if cnt <= k {
		io.Println("Yes")
	} else {
		io.Println("No")
	}
}
