package main

import (
	"math"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	x, y, u = make([]int, n), make([]int, n), make([]string, n)
	for i := 0; i < n; i++ {
		x[i], y[i], u[i] = io.NextInt(), io.NextInt(), io.Next()
	}
	waU, waD, waL, waR := map[int][]int{}, map[int][]int{}, map[int][]int{}, map[int][]int{}
	saU, saD, saL, saR := map[int][]int{}, map[int][]int{}, map[int][]int{}, map[int][]int{}
	min := inf
	for i := range u {
		wa, sa := x[i]+y[i], x[i]-y[i]
		var v0, v1 int
		switch u[i] {
		case "U":
			waU[wa] = append(waU[wa], i)
			saU[sa] = append(saU[sa], i)
		case "D":
			waD[wa] = append(waD[wa], i)
			saD[sa] = append(saD[wa], i)
		case "L":
			waL[wa] = append(waL[wa], i)
			saL[sa] = append(saL[sa], i)
		case "R":
			waR[wa] = append(waR[wa], i)
			saR[sa] = append(saR[sa], i)
		}
		if v0 < min {
			min = v0
		}
		if v1 < min {
			min = v1
		}
	}
}

const inf = math.MaxInt64

var x, y []int
var u []string

func vCollisionTime(key int, mp map[int][]int) int {
	return 0
}
