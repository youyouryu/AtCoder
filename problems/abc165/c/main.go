package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

var n, m, q int
var a, b, c, d []int

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, m, q = io.NextInt(), io.NextInt(), io.NextInt()
	a, b, c, d = make([]int, q), make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		a[i], b[i], c[i], d[i] = io.NextInt(), io.NextInt(), io.NextInt(), io.NextInt()
	}
	ans := solve()
	io.Println(ans)
}

func solve() (ans int) {
	return search([]int{1})
}

func search(choice []int) (max int) {
	if len(choice) == n+1 {
		return scoring(choice)
	}
	for i := choice[len(choice)-1]; i <= m; i++ {
		score := search(append(choice, i))
		if score > max {
			max = score
		}
	}
	return max
}

func scoring(choice []int) int {
	score := 0
	for i := range a {
		if choice[b[i]]-choice[a[i]] == c[i] {
			score += d[i]
		}
	}
	return score
}
