package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	x, k, d := io.NextInt(), io.NextInt(), io.NextInt()
	x = lib.Abs(x)
	q := x / d
	r := x % d
	if k-q < 0 {
		x = x - k*d
	} else if (k-q)%2 == 0 {
		x = r
	} else {
		x = lib.Abs(r - d)
	}
	io.Println(x)
}
