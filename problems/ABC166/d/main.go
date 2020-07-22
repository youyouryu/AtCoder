package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	x := io.NextInt()
	a, b := solve(x)
	io.Println(a, b)
}

func solve(x int) (int, int) {
	for a := 1; pow(a, 5)-pow(a-1, 5) <= x; a++ {
		for b := a - 1; pow(a, 5)-pow(b, 5) <= x; b-- {
			if pow(a, 5)-pow(b, 5) == x {
				return a, b
			}
		}
	}
	return 0, 0
}

func pow(base, n int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret *= base
		}
		base *= base
		n >>= 1
	}
	return ret
}
