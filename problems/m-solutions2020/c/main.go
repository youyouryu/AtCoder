package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, k := io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	for i := k; i < n; i++ {
		if a[i] > a[i-k] {
			io.Println("Yes")
		} else {
			io.Println("No")
		}
	}
}
