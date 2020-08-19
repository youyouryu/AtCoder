package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	s := io.Next()
	cnt, max := 0, 0
	for i := range s {
		if s[i] == 'R' {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
		}
	}
	if cnt > max {
		max = cnt
	}
	io.Println(max)
}
