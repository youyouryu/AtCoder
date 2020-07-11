package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	io := NewIo(os.Stdin, os.Stdout)
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

// Io is I/O object
type Io struct {
	reader    *bufio.Reader
	writer    *bufio.Writer
	tokens    []string
	nextToken int
}

// NewIo returns a new Io instance
func NewIo(r io.Reader, w io.Writer) *Io {
	return &Io{
		reader: bufio.NewReader(r),
		writer: bufio.NewWriter(w),
	}
}

// Flush flushes writer
func (io *Io) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}

// NextLine scans a line from stdin
func (io *Io) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

// Next return a word from stdin
func (io *Io) Next() string {
	for io.nextToken >= len(io.tokens) {
		line := io.NextLine()
		io.tokens = strings.Fields(line)
		io.nextToken = 0
	}
	r := io.tokens[io.nextToken]
	io.nextToken++
	return r
}

// NextInt returns an integer from stdin
func (io *Io) NextInt() int {
	i, err := strconv.Atoi(io.Next())
	if err != nil {
		panic(err)
	}
	return i
}

// Println is a wrapper of fmt.Fprintln
func (io *Io) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}
