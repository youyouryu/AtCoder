package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	h, w := io.NextInt(), io.NextInt()
	stage := make([]string, h)
	for i := 0; i < h; i++ {
		stage[i] = io.Next()
	}
	ans := solve(h, w, stage)
	io.Println(ans)
}

func solve(h, w int, stage []string) (ans int) {
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	for x := 0; x < h; x++ {
		for y := 0; y < w; y++ {
			if visited[x][y] {
				continue
			}
			ans += bfs(x, y, visited, stage)
		}
	}
	return ans
}

func bfs(x, y int, visited [][]bool, stage []string) int {
	h, w := len(stage), len(stage[0])
	count := make(map[byte]int, 2)
	q := newQueue()
	visited[x][y] = true
	count[stage[x][y]]++
	q.push(pair{x, y})
	for len(*q) > 0 {
		p := q.pop()
		for i := range dx {
			nx, ny := p.x+dx[i], p.y+dy[i]
			if nx < 0 || h <= nx || ny < 0 || w <= ny {
				continue
			}
			if visited[nx][ny] {
				continue
			}
			if stage[nx][ny] == stage[p.x][p.y] {
				continue
			}
			visited[nx][ny] = true
			count[stage[nx][ny]]++
			q.push(pair{nx, ny})
		}
	}
	return count['#'] * count['.']
}

type pair struct{ x, y int }

type queue []pair

func newQueue() *queue {
	return &queue{}
}

func (q *queue) push(p pair) {
	*q = append(*q, p)
}

func (q *queue) pop() pair {
	p := (*q)[0]
	*q = (*q)[1:]
	return p
}

var dx = []int{0, 0, -1, 1}
var dy = []int{-1, 1, 0, 0}
