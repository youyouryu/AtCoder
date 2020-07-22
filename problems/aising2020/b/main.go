package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	a := io.NextInts(n)
	count := 0
	for i := 0; i < n; i += 2 {
		if a[i]%2 == 1 {
			count++
		}
	}
	io.Println(count)
}
