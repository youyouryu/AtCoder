package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	x := io.NextInt()
	y := 100
	i := 0
	for y < x {
		y += y / 100
		i++
	}
	io.Println(i)
}
