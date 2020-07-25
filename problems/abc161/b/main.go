package main

import (
	"os"
	"sort"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, m := io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	sort.Ints(a)
	sum := 0
	for i := range a {
		sum += a[i]
	}
	for i := 0; i < m; i++ {
		if a[n-1-i]*4*m < sum {
			io.Println("No")
			return
		}
	}
	io.Println("Yes")
}
