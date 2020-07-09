package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = sc.NextInt()
		b[i], _ = sc.NextInt()
	}
	ans := solve(n, a, b)
	fmt.Println(ans)
}

const mod = 1000000007

type pair struct {
	x, y int
}

func solve(n int, a, b []int) (ans uint) {
	zero := uint(0)
	count := map[pair]*pair{}
	for i := 0; i < n; i++ {
		x, y := a[i], b[i]
		if x == 0 && y == 0 {
			zero++
			continue
		}
		if y < 0 || y == 0 && x < 0 {
			x, y = rotate180(x, y)
		}
		rot90 := x <= 0
		if rot90 {
			x, y = rotate90(x, y)
		}
		g := gcd(x, y)
		x, y = x/g, y/g
		key := pair{x, y}
		if count[key] == nil {
			count[key] = &pair{}
		}
		if rot90 {
			count[key].y++
		} else {
			count[key].x++
		}
	}

	ans = 1
	for _, v := range count {
		c1, c2 := v.x, v.y
		now := uint(1)
		now = (now + lib.ModPow(2, uint(c1), mod) - 1) % mod
		now = (now + lib.ModPow(2, uint(c2), mod) - 1) % mod
		ans = lib.ModMul(ans, now, mod)
	}
	if ans == 0 {
		panic("")
	}
	ans = (ans + zero - 1 + mod) % mod
	return ans
}

func rotate180(x1, y1 int) (x2, y2 int) {
	return -x1, -y1
}

func rotate90(x1, y1 int) (x2, y2 int) {
	return y1, -x1
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
