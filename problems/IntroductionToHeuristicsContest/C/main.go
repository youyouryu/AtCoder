package main

import (
	"os"

	"github.com/yuyamada/AtCoder/lib"
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
	m := io.NextInt()
	k, v := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		k[i], v[i] = io.NextInt(), io.NextInt()
	}
	ans := solve(c, s, t, k, v)
	for i := range ans {
		io.Println(ans[i])
	}
}

func solve(c []int, s [][]int, t, k, v []int) []int {
	m := len(k)
	scores := make([]int, m)
	for i := 0; i < m; i++ {
		t[k[i]-1] = v[i]
		scores[i] = calcScore(c, s, t)
	}
	return scores
}

func calcScore(c []int, s [][]int, t []int) int {
	d := len(s)
	lasts := make([]int, n)
	score := 0
	for i := 0; i < d; i++ {
		lasts[t[i]-1] = i + 1
		up := s[i][t[i]-1]
		down := calcDown(c, lasts, i+1)
		score += up - down
	}
	return score
}

func calcDown(c []int, lasts []int, day int) int {
	down := 0
	for i := 0; i < n; i++ {
		down += c[i] * (day - lasts[i])
	}
	return down
}
