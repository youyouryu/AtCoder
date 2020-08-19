package main

import (
	"math"
	"os"
	"sort"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, k := io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	ans := solve(n, k, a)
	io.Println(ans)
}

func solve(n, k int, a []int) int {
	divConut := func(x float64) int {
		if x == 0 {
			panic("division by zero")
		}
		cnt := 0
		for i := range a {
			if float64(a[i]) <= x {
				continue
			}
			cnt += int(math.Ceil(float64(a[i])/x)) - 1
		}
		return cnt
	}
	x := sort.Search(maxIter, func(x int) bool { return divConut(float64(x+1)/10) <= k })
	if x == maxIter {
		panic("reached max iteration")
	}
	return int(math.Ceil(float64(x+1) / 10))
}

const maxIter = 1e12
