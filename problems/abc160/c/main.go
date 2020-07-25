package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	k, n := io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	max := k + a[0] - a[n-1]
	for i := 0; i < n-1; i++ {
		if a[i+1]-a[i] > max {
			max = a[i+1] - a[i]
		}
	}
	io.Println(k - max)
}
