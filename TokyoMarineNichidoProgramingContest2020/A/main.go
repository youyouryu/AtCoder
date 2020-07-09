package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	fmt.Println(string([]rune(s)[:3]))
}
