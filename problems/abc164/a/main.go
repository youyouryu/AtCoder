package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	s, w := io.NextInt(), io.NextInt()
	if s <= w {
		io.Println("unsafe")
	} else {
		io.Println("safe")
	}
}
