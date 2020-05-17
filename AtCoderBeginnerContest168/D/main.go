package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type vertex struct {
	id   int
	dist int
}

type priorityQueue []vertex

func (p priorityQueue) Len() int            { return len(p) }
func (p priorityQueue) Less(i, j int) bool  { return p[i].dist < p[j].dist }
func (p priorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{}) { *p = append(*p, x.(vertex)) }
func (p *priorityQueue) Pop() interface{} {
	old := *p
	*p = old[:len(old)-1]
	return old[len(old)-1]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(line[0])
	m, _ := strconv.Atoi(line[1])
	graph := make([][]int, n)
	for i := 0; i < m; i++ {
		sc.Scan()
		line = strings.Split(sc.Text(), " ")
		from, _ := strconv.Atoi(line[0])
		to, _ := strconv.Atoi(line[1])
		graph[from-1] = append(graph[from-1], to-1)
		graph[to-1] = append(graph[to-1], from-1)
	}

	dists := dikstra(&graph)
	if route, ok := traceRoute(&graph, dists); ok {
		fmt.Println("Yes")
		for i := range *route {
			fmt.Println((*route)[i])
		}
	} else {
		fmt.Println("No")
	}
}

func dikstra(graph *[][]int) *[]int {
	dists := make([]int, len(*graph))
	for i := range dists {
		dists[i] = math.MaxInt64
	}
	dists[0] = 0

	queue := priorityQueue{}
	heap.Push(&queue, vertex{id: 0, dist: 0})
	for len(queue) > 0 {
		v := heap.Pop(&queue).(vertex)
		d := v.dist + 1
		for _, id := range (*graph)[v.id] {
			if dists[id] != math.MaxInt64 {
				continue
			}
			if d < dists[id] {
				dists[id] = d
				heap.Push(&queue, vertex{id: id, dist: d})
			}
		}
	}
	return &dists
}

func traceRoute(graph *[][]int, dists *[]int) (route *[]int, ok bool) {
	route = &[]int{}
	for i := range *graph {
		if i == 0 {
			*route = append(*route, 0)
			continue
		}
		var min, argmin int
		for j, id := range (*graph)[i] {
			if j == 0 || (*dists)[id] < min {
				min, argmin = (*dists)[id], id
			}
		}
		if argmin == math.MaxInt64 {
			return nil, false
		}
		*route = append(*route, argmin+1)
	}
	return route, true
}
