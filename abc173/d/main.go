package main

import (
	"os"
	"sort"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	a := io.NextInts(n)
	ans := solve(a)
	io.Println(ans)
}

func solve(a []int) (ans int) {
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	q := []int{a[0]}
	for i := 1; i < len(a); i++ {
		ans += q[0]
		q = append(q[1:], []int{a[i], a[i]}...)
	}
	return ans
}
