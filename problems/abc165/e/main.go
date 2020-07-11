package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, m := io.NextInt(), io.NextInt()
	a, b := make([]int, m), make([]int, m)
	i, j := 1, 1
	idx := 0
	if m%2 == 1 {
		a[idx] = i
		b[idx] = n + 1 - i
		i++
		idx++
	}
	for idx < m {
		a[idx] = i
		b[idx] = n + 1 - i
		i++
		idx++
		a[idx] = n/2 + 1 - j
		b[idx] = n/2 + 1 + j
		j++
		idx++
	}

	for i := range a {
		io.Println(a[i], b[i])
	}
}
