package main

import (
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

const n = 26

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	d := io.NextInt()
	c := io.NextInts(n)
	s := make([][]int, d)
	for i := 0; i < d; i++ {
		s[i] = io.NextInts(n)
	}
	t := io.NextInts(d)
	ans := solve(c, s, t)
	for i := range ans {
		io.Println(ans[i])
	}
}

func solve(c []int, s [][]int, t []int) []int {
	d := len(s)
	lasts := make([]int, n)
	scores := make([]int, d+1)
	for i := 0; i < d; i++ {
		lasts[t[i]-1] = i + 1
		up := s[i][t[i]-1]
		down := calcDown(c, lasts, i+1)
		scores[i+1] = scores[i] + up - down
	}
	return scores[1:]
}

func calcDown(c []int, lasts []int, day int) int {
	down := 0
	for i := 0; i < n; i++ {
		down += c[i] * (day - lasts[i])
	}
	return down
}
