package main

import (
	"math"
	"os"
	"sort"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	a := io.NextInts(n)
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		ui, vi := io.NextInt(), io.NextInt()
		graph[ui-1] = append(graph[ui-1], vi-1)
		graph[vi-1] = append(graph[vi-1], ui-1)
	}
	ans := solve(n, a, graph)
	for i := range ans {
		io.Println(ans[i])
	}
}

func solve(n int, a []int, graph [][]int) []int {
	stk := newStack()
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = inf
	}
	lis := make([]int, n)
	stk.push(pair{0, inf})
	dfs(0, stk, dp, lis, graph, a)
	return lis
}

func dfs(from int, stk *stack, dp []int, lis []int, graph [][]int, a []int) {
	if lis[from] != 0 {
		return
	}
	i := sort.Search(len(dp), func(i int) bool { return dp[i] >= a[from] })
	stk.push(pair{i, dp[i]})
	dp[i] = a[from]
	j := sort.Search(len(dp), func(j int) bool { return dp[j] >= inf })
	lis[from] = j
	for _, to := range graph[from] {
		dfs(to, stk, dp, lis, graph, a)
	}
	top := stk.pop()
	dp[top.key] = top.value
}

const inf = math.MaxInt64

type pair struct{ key, value int }

type stack []pair

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(x pair) {
	*s = append(*s, x)
}

func (s *stack) pop() pair {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}
