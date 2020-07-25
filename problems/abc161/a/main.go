package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	a, b, c := io.NextInt(), io.NextInt(), io.NextInt()
	a, b = b, a
	a, c = c, a
	io.Println(a, b, c)
}
