package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	k := io.NextInt()
	a := make([][]int, k)
	for i := 0; i < k; i++ {
		di := io.NextInt()
		a[i] = io.NextInts(di)
	}
	io.Println("hello")
	io.Println(n, k)
	for i := 0; i < k; i++ {
		io.Println(a[i])
	}
}
