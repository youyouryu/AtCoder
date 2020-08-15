package main

import (
	"math"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, d := io.NextInt(), io.NextFloat()
	cnt := 0
	for i := 0; i < n; i++ {
		xi, yi := io.NextFloat(), io.NextFloat()
		dist := math.Sqrt(xi*xi + yi*yi)
		if dist <= d {
			cnt++
		}
	}
	io.Println(cnt)
}
