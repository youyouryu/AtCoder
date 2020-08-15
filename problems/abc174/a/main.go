package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	x := io.NextInt()
	if x >= 30 {
		io.Println("Yes")
	} else {
		io.Println("No")
	}
}
