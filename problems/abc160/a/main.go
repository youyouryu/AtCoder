package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	r := []rune(io.Next())
	if r[2] == r[3] && r[4] == r[5] {
		io.Println("Yes")
	} else {
		io.Println("No")
	}
}
