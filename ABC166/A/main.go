package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	s, _ := sc.NextString()
	if s == "ABC" {
		fmt.Println("ARC")
	} else {
		fmt.Println("ABC")
	}
}
