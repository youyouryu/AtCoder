package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	count := map[string]int{}
	for i := 0; i < n; i++ {
		count[io.Next()]++
	}
	io.Println(len(count))
}
