package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	a, b, c, d := io.NextInt(), io.NextInt(), io.NextInt(), io.NextInt()
	var out string
	for {
		c -= b
		if c <= 0 {
			out = "Yes"
			break
		}
		a -= d
		if a <= 0 {
			out = "No"
			break
		}
	}
	io.Println(out)
}
