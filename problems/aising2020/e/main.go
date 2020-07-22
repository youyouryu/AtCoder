package main

import (
	"container/heap"
	"os"
	"sort"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	t := io.NextInt()
	tcs := make([]testCase, t)
	for i := 0; i < t; i++ {
		n := io.NextInt()
		k, l, r := make([]int, n), make([]int, n), make([]int, n)
		for j := 0; j < n; j++ {
			k[j], l[j], r[j] = io.NextInt(), io.NextInt(), io.NextInt()
		}
		tcs[i] = testCase{n, k, l, r}
	}
	// O(t n log n) = O(2e10 log 2e5)
	for i := range tcs {
		io.Println(solve(tcs[i]))
	}
}

// O(n log n)
func solve(tc testCase) (ans int) {
	n, k, l, r := tc.n, tc.k, tc.l, tc.r
	g0, g1 := []int{}, []int{}
	for i := 0; i < n; i++ {
		if l[i] >= r[i] {
			g0 = append(g0, i)
		} else {
			g1 = append(g1, i)
		}
	}
	sort.Slice(g0, func(i, j int) bool { return k[g0[i]] < k[g0[j]] })
	sort.Slice(g1, func(i, j int) bool { return k[g1[i]] > k[g1[j]] })
	ans += sortAndScore(g0, n, k, l, r)
	kInv := make([]int, n)
	for _, idx := range g1 {
		kInv[idx] = n - k[idx]
	}
	ans += sortAndScore(g1, n, kInv, r, l)
	return
}

func sortAndScore(g []int, n int, k, l, r []int) (ans int) {
	for _, idx := range g {
		ans += r[idx]
	}
	q := lib.IntHeap{}
	for i := 1; i <= n; i++ {
		begin := sort.Search(len(g), func(j int) bool { return k[g[j]] >= i })
		end := sort.Search(len(g), func(j int) bool { return k[g[j]] > i })
		if begin == end {
			continue
		}
		for j := begin; j < end; j++ {
			idx := g[j]
			heap.Push(&q, l[idx]-r[idx])
		}
		for q.Len() > i {
			heap.Pop(&q)
		}
	}
	for i := range q {
		ans += q[i]
	}
	return ans
}

type testCase struct {
	n       int
	k, l, r []int
}
