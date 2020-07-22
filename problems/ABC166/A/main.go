package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	s := io.Next()
	if s == "ABC" {
		io.Println("ARC")
	} else {
		io.Println("ABC")
	}
}
