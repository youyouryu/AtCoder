package main

import (
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	n %= 1000
	if n != 0 {
		n = 1000 - n
	}
	io.Println(n)
}
