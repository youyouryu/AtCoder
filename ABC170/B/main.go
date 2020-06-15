package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	x, _ := sc.NextInt()
	y, _ := sc.NextInt()
	ans := solve(x, y)
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func solve(x, y int) bool {
	i := sort.Search(x+1, func(i int) bool { return 4*i+2*(x-i) >= y })
	return i < x+1 && 4*i+2*(x-i) == y
}
