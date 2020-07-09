package main

import (
	"os"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	k := io.NextInt()
	a := io.NextInt()
	b := io.NextInt()
	for i := a; i <= b; i++ {
		if i%k == 0 {
			io.Println("OK")
			return
		}
	}
	io.Println("NG")
}
