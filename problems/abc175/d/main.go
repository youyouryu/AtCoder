package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, k := io.NextInt(), io.NextInt()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = io.NextInt() - 1
	}
	c := io.NextInts(n)
	ans := solve(n, k, p, c)
	io.Println(ans)
}

func solve(n, k int, p, c []int) int {
	var max int
	for i := 0; i < n; i++ {
		score := play(n, k, i, p, c)
		if i == 0 || score > max {
			max = score
		}
	}
	return max
}

func play(n, k, i int, p, c []int) int {
	score := 0
	scores := []int{}
	j := i
	for {
		score += c[p[j]]
		scores = append(scores, score)
		j = p[j]
		if j == i {
			break
		}
	}
	if k <= len(scores) {
		return lib.Max(scores[:k]...)
	}
	if score <= 0 {
		return lib.Max(scores[:lib.Min(k, len(scores))]...)
	}
	q := k / len(scores)
	r := k % len(scores)
	s1 := lib.Max(scores...)
	s2 := score
	if r != 0 {
		s2 += lib.Max(0, lib.Max(scores[:r]...))
	}
	score *= q - 1
	score += lib.Max(s1, s2)
	return score
}
