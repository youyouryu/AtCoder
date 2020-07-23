package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	ans := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			continue
		}
		ans += i
	}
	io.Println(ans)
}
