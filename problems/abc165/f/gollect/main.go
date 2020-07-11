package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	io := NewIo(os.Stdin, os.Stdout)
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

// NextInts returns n integers from stdin
func (io *Io) NextInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = io.NextInt()
	}
	return ret
}

// Println is a wrapper of fmt.Fprintln
func (io *Io) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}
