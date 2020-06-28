package main

import (
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	s, t := io.Next(), io.Next()
	count := 0
	for i := range s {
		if s[i] != t[i] {
			count++
		}
	}
	io.Println(count)
}
