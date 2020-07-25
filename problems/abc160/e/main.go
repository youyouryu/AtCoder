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
	x, y, a, b, c := io.NextInt(), io.NextInt(), io.NextInt(), io.NextInt(), io.NextInt()
	p, q, r := io.NextInts(a), io.NextInts(b), io.NextInts(c)
	sort.Ints(p)
	sort.Ints(q)
	sort.Ints(r)
	h := &lib.IntHeap{}
	for i := 0; i < x; i++ {
		heap.Push(h, p[a-1-i])
	}
	for i := 0; i < y; i++ {
		heap.Push(h, q[b-1-i])
	}
	for i := range r {
		heap.Push(h, r[c-1-i])
		heap.Pop(h)
	}
	ans := 0
	for h.Len() > 0 {
		ans += heap.Pop(h).(int)
	}
	io.Println(ans)
}
