package main

import (
	"fmt"
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	a, _ := sc.NextInt()
	v, _ := sc.NextInt()
	b, _ := sc.NextInt()
	w, _ := sc.NextInt()
	t, _ := sc.NextInt()
	ans := solve(a, v, b, w, t)
	if ans {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(a, v, b, w, t int) bool {
	dist := a - b
	if dist < 0 {
		dist *= -1
	}
	velocity := v - w
	return dist-velocity*t <= 0
}
