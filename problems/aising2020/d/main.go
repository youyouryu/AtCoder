package main

import (
	"log"
	"math/bits"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

var logger = log.New(os.Stderr, "[DEBUG] ", 0)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	_ = io.NextInt()
	s := io.Next()
	ans := solve2(s)
	for _, v := range ans {
		io.Println(v)
	}
}

func solve(s string) []int {
	n := len(s)
	x := make([]int, n)
	for i := range s {
		x[i] = int(s[i] - '0')
	}
	pc0 := 0
	for i := range x {
		pc0 += x[i]
	}
	r0 := 0
	for i := 0; i < n; i++ {
		r0 = (r0 * 2) % pc0
		r0 += x[i]
	}
	ans := make([]int, n)
	dr := 1
	for i := n - 1; i >= 0; i-- {
		pc, r := pc0, r0
		if x[i] == 0 {
			pc++
			r = (r + dr) % pc
		} else {
			pc--
			r = (r - dr + pc) % pc
		}
		// pcが変わるのでdrを再利用できない
		dr = (dr * 2) % pc
		ans[i] = f(r) + 1
	}

	return ans
}

func solve2(s string) []int {
	n := len(s)
	x := make([]int, n)
	for i := range s {
		x[i] = int(s[i] - '0')
	}
	pc0 := 0
	for i := range x {
		pc0 += x[i]
	}

	ans := make([]int, n)
	// 桁の値でforを回す。
	// 2n回ループを回すが、if文に引っかかるのはn回のみ。
	// 桁の値が同じものをまとめて処理することで、dr%pcを再利用できる
	for b := 0; b < 2; b++ {
		pc := pc0
		if b == 0 {
			pc++
		} else {
			pc--
		}
		if pc <= 0 {
			continue
		}
		r0 := 0
		for i := 0; i < n; i++ {
			r0 = (r0 * 2) % pc
			r0 += x[i]
		}
		dr := 1
		for i := n - 1; i >= 0; i-- {
			r := r0
			if x[i] == b {
				if b == 0 {
					r = (r + dr) % pc
				} else {
					r = (r - dr + pc) % pc
				}
				ans[i] = f(r) + 1
			}
			// 0->1を一括で処理
			// 1->0を一括で処理
			dr = (dr * 2) % pc
		}
	}
	return ans
}

func f(x int) int {
	if x == 0 {
		return 0
	}
	return f(x%pcnt(x)) + 1
}

func pcnt(x int) int {
	return bits.OnesCount(uint(x))
}
