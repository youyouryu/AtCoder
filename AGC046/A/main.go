package main

import (
	"fmt"
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	x, _ := sc.NextInt()
	ans := solve(x)
	fmt.Println(ans)
}

func solve(x int) int {
	a, b := 360, x
	return a / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
