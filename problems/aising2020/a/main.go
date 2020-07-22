package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	l, r, d := io.NextInt(), io.NextInt(), io.NextInt()
	count := 0
	for i := l; i <= r; i++ {
		if i%d == 0 {
			count++
		}
	}
	io.Println(count)
}
