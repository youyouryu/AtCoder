package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	s := io.Next()
	for i := range s {
		if s[i] == '7' {
			io.Println("Yes")
			return
		}
	}
	io.Println("No")
}
