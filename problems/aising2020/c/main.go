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
	n := io.NextInt()
	for k := 1; k <= n; k++ {
		maxIter := int(math.Sqrt(float64(k)))
		count := 0
		for x := 1; x <= maxIter; x++ {
			for y := 1; y <= maxIter; y++ {
				z := sort.Search(maxIter, func(z int) bool { return g(x, y, z) >= k })
				if z != 0 && g(x, y, z) == k {
					count++
				}
			}
		}
		io.Println(count)
	}
}

func g(x, y, z int) int {
	return (x+y+z)*(x+y+z) - x*y - y*z - z*x
}
