package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	h, _ := sc.NextInt()
	w, _ := sc.NextInt()
	k, _ := sc.NextInt()
	x1, _ := sc.NextInt()
	y1, _ := sc.NextInt()
	x2, _ := sc.NextInt()
	y2, _ := sc.NextInt()
	stage := make([]string, h)
	for i := 0; i < h; i++ {
		sc.Scan()
		stage[i] = sc.Text()
	}
	ans := solve(h, w, k, x1-1, y1-1, x2-1, y2-1, stage)
	fmt.Println(ans)
}

type state struct {
	x, y, dist int
}

// gollect: keep methods
type priorityQueue []*state

func (p priorityQueue) Len() int            { return len(p) }
func (p priorityQueue) Less(i, j int) bool  { return p[i].dist < p[j].dist }
func (p priorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{}) { *p = append(*p, x.(*state)) }
func (p *priorityQueue) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

var inf = math.MaxInt64
var dxs, dys = []int{0, 0, -1, 1}, []int{-1, 1, 0, 0}

func solve(h, w, k, x1, y1, x2, y2 int, stage []string) int {
	dists := make([][]int, h)
	for x := range dists {
		dists[x] = make([]int, w)
		for y := range dists[x] {
			dists[x][y] = inf
		}
	}

	dists[x1][y1] = 0
	q := priorityQueue{}
	s := &state{x1, y1, 0}
	heap.Push(&q, s)
	for len(q) > 0 {
		p := heap.Pop(&q).(*state)
		if p.x == x2 && p.y == y2 {
			return p.dist
		}
		for d := range dxs {
			for i := 1; i <= k; i++ {
				x, y := p.x+i*dxs[d], p.y+i*dys[d]
				if x < 0 || h <= x || y < 0 || w <= y {
					break
				}
				if stage[x][y] == '@' {
					break
				}
				dist := p.dist + 1
				if dist > dists[x][y] {
					break
				}
				if dist == dists[x][y] {
					continue
				}
				dists[x][y] = dist
				s := &state{x, y, dist}
				heap.Push(&q, s)
			}

		}
	}
	if dists[x2][y2] == inf {
		return -1
	}
	return dists[x2][y2]
}
