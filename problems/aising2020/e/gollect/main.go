package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	io := NewIo(os.Stdin, os.Stdout)
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
	q := IntHeap{}
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

// IntHeap is heap for interger

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to heap
func (h *IntHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}

// Pop returns the minimum element of the heap and remove it
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
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

// Println is a wrapper of fmt.Fprintln
func (io *Io) Println(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}
