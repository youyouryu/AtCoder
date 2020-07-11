package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, h, w := io.NextInt(), io.NextInt(), io.NextInt()
	count := 0
	for i := 0; i+h <= n; i++ {
		for j := 0; j+w <= n; j++ {
			count++
		}
	}
	io.Println(count)
}
