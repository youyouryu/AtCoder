package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	a, _ := sc.NextInts(5)
	for i, v := range a {
		if v == 0 {
			fmt.Print(i + 1)
			return
		}
	}
}
