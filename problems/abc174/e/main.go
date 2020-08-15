package main

import (
	"container/heap"
	"math"
	"os"
	"sort"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, k := io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	ans := solve(n, k, a)
	io.Println(ans)
}

func solve(n, k int, a []int) int {
	if n == 1 {
		return int(math.Ceil(float64(a[0]) / float64(k)))

	}
	q := &lib.PriorityQueue{}
	for i := range a {
		heap.Push(q, wood{float64(a[i]), 1})
	}
	for i := 0; i < k; i++ {
		w0 := heap.Pop(q).(wood)
		w1 := heap.Pop(q).(wood)
		j := sort.Search(k-i, func(j int) bool { return w0.length/(w0.division+float64(j)) <= w1.length/w1.division })
		w0.division += float64(j)
		i += j - 1
		heap.Push(q, w0)
		heap.Push(q, w1)
	}
	w := heap.Pop(q).(wood)
	return int(math.Ceil(w.length / w.division))
}

// gollect: keep methods
type wood struct {
	length   float64
	division float64
}

func (w wood) Less(z lib.Comparable) bool {
	return w.length/w.division > z.(wood).length/z.(wood).division
}
