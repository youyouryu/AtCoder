package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	fmt.Println(string([]rune(s)[:3]))
}
