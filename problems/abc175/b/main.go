package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	l := io.NextInts(n)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if l[i] == l[j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if l[j] == l[k] || l[k] == l[i] {
					continue
				}
				l1, l2, l3 := l[i], l[j], l[k]
				l1, l2 = swap(l1, l2)
				l2, l3 = swap(l2, l3)
				l1, l2 = swap(l1, l2)
				if l1+l2 > l3 {
					cnt++
				}
			}
		}
	}
	io.Println(cnt)
}

func swap(i, j int) (int, int) {
	if j < i {
		return j, i
	}
	return i, j
}
