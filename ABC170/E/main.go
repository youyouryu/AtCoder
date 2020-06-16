package main

import (
	"container/heap"
	"fmt"
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	q, _ := sc.NextInt()
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		ai, _ := sc.NextInt()
		bi, _ := sc.NextInt()
		a[i], b[i] = ai, bi
	}
	c, d := make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		ci, _ := sc.NextInt()
		di, _ := sc.NextInt()
		c[i], d[i] = ci, di
	}
	ans := solve(a, b, c, d)
	for i := range ans {
		fmt.Println(ans[i])
	}
}

type infant struct {
	id, rate, belongs int
}

// gollect: keep methods
type priorityQueue []infant

func (p priorityQueue) Len() int            { return len(p) }
func (p priorityQueue) Less(i, j int) bool  { return p[i].rate < p[j].rate }
func (p priorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{}) { *p = append(*p, x.(infant)) }
func (p *priorityQueue) Pop() interface{} {
	old := *p
	*p = old[:len(old)-1]
	return old[len(old)-1]
}

func solve(a, b, c, d []int) []int {
	pqs := make([]priorityQueue, 2e5)
	for i := range a {
		inf := infant{id: i + 1, rate: -a[i], belongs: b[i]}
		heap.Push(&pqs[b[i]-1], inf)
	}
	pqMax := priorityQueue{}
	for i := range pqs {
		update(&pqMax, &pqs[i], b)
	}

	ans := make([]int, len(c))
	for i := range c {
		src := b[c[i]-1]
		b[c[i]-1] = d[i]
		update(&pqMax, &pqs[src-1], b)
		heap.Push(&pqs[d[i]-1], infant{id: c[i], rate: -a[c[i]-1], belongs: d[i]})
		update(&pqMax, &pqs[d[i]-1], b)
		ans[i] = equality(&pqMax, pqs)
	}
	return ans
}

func update(pqMax, pq *priorityQueue, belongs []int) {
	for len(*pq) > 0 {
		inf := (*pq)[0]
		if inf.belongs == belongs[inf.id-1] {
			heap.Push(pqMax, infant{inf.id, -inf.rate, inf.belongs})
			break
		}
		heap.Pop(pq)
	}
}

func equality(pqMax *priorityQueue, pqs []priorityQueue) int {
	for len(*pqMax) > 0 {
		inf := (*pqMax)[0]
		pq := pqs[inf.belongs-1]
		if len(pq) > 0 && inf.id == pq[0].id {
			return inf.rate
		}
		heap.Pop(pqMax)
	}
	return -1
}
