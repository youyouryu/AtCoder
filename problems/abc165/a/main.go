package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	k, a, b := io.NextInt(), io.NextInt(), io.NextInt()
	for i := a; i <= b; i++ {
		if i%k == 0 {
			io.Println("OK")
			return
		}
	}
	io.Println("NG")
}
