package main

import (
	"math"
	"os"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	d = io.NextInt()
	c := io.NextInts(n)
	s := make([][]int, d)
	for i := 0; i < d; i++ {
		s[i] = io.NextInts(n)
	}
	ans := solve(c, s)
	for i := range ans {
		io.Println(ans[i])
	}
}

const (
	n    = 26
	ninf = -math.MaxInt64
)

var d int

func solve(c []int, s [][]int) []int {
	return greedy(c, s)
}

func greedy(c []int, s [][]int) []int {
	t := make([]int, d)
	lasts := make([]int, n)
	for i := 0; i < d; i++ {
		if i >= len(s) {
			break
		}
		_, ti := greedyStep(c, s, lasts, i+1)
		t[i] = ti
		lasts[ti-1] = i + 1
	}
	return t[:len(s)]
}

func greedyStep(c []int, s [][]int, lasts []int, day int) (int, int) {
	max, argmax := ninf, -1
	for i := 0; i < n; i++ {
		delta := calcDelta(c, s, lasts, day, i+1)
		if delta > max {
			max, argmax = delta, i+1
		}
	}
	return max, argmax
}

func calcDelta(c []int, s [][]int, lasts []int, day, contest int) int {
	up := s[day-1][contest-1]
	lasts[contest-1] = day
	down := calcDown(c, lasts, day)
	return up - down
}

func calcDown(c []int, lasts []int, day int) int {
	down := 0
	for i := 0; i < n; i++ {
		down += c[i] * (day - lasts[i])
	}
	return down
}

func hello() string {
	return "hello"
}
