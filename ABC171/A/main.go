package main

import (
	"fmt"
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	s, _ := sc.NextString()
	a := []rune(s)[0]
	var out string
	if 'a' <= a && a <= 'z' {
		out = "a"
	} else {
		out = "A"
	}
	fmt.Println(out)
}
