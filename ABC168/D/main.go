package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	m, _ := sc.NextInt()
	a, b := []int{}, []int{}
	for i := 0; i < m; i++ {
		ai, _ := sc.NextInt()
		bi, _ := sc.NextInt()
		a = append(a, ai)
		b = append(b, bi)
	}
	ans := solve(n, a, b)
	if ans == nil {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
		for i := range ans {
			fmt.Println(ans[i])
		}

	}
}

func solve(n int, a, b []int) (ans []int) {
	cl := createConcatenatedList(n, a, b)
	dists := dijkstra(n, cl)
	labels := toLabels(cl, dists)
	return labels
}

func createConcatenatedList(n int, a, b []int) [][]int {
	cl := make([][]int, n)
	for i := range a {
		cl[a[i]-1] = append(cl[a[i]-1], b[i])
		cl[b[i]-1] = append(cl[b[i]-1], a[i])
	}
	return cl
}

func dijkstra(n int, cl [][]int) (minDists []int) {
	dists := createDistanceList(n)
	q := lib.PriorityQueue{}
	heap.Push(&q, lib.Pair{Key: 1, Value: 0})
	for len(q) > 0 {
		from := heap.Pop(&q).(lib.Pair)
		dist := from.Value + 1
		for _, key := range cl[from.Key-1] {
			if dist < dists[key-1] {
				dists[key-1] = dist
				to := lib.Pair{Key: key, Value: dist}
				heap.Push(&q, to)
			}
		}
	}
	return dists
}

func createDistanceList(n int) []int {
	dists := make([]int, n)
	for i := range dists {
		dists[i] = math.MaxInt64
	}
	dists[0] = 0
	return dists
}

func toLabels(cl [][]int, minDists []int) []int {
	labels := []int{}
	for i := 1; i < len(cl); i++ {
		var argmin, min int
		for j := range cl[i] {
			if j == 0 || minDists[cl[i][j]-1] < min {
				argmin, min = cl[i][j], minDists[cl[i][j]-1]
			}
		}
		labels = append(labels, argmin)
	}
	return labels
}
