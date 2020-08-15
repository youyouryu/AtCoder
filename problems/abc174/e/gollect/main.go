package main

import (
	"bufio"
	"container/heap"
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
	n, k := io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	ans := solve(n, k, a)
	io.Println(ans)
}

func solve(n, k int, a []int) int {
	if n == 1 {
		return int(math.Ceil(float64(a[0]) / float64(k)))

	}
	q := &PriorityQueue{}
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

type wood struct {
	length   float64
	division float64
}

func (w wood) Less(z Comparable) bool {
	return w.length/w.division > z.(wood).length/z.(wood).division
}

type Comparable interface{ Less(Comparable) bool }

// PriorityQueue

type PriorityQueue []Comparable

func (q PriorityQueue) Len() int           { return len(q) }
func (q PriorityQueue) Less(i, j int) bool { return q[i].Less(q[j]) }
func (q PriorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

// Push adds an element
func (q *PriorityQueue) Push(x interface{}) { *q = append(*q, x.(Comparable)) }

// Pop gets and remove the minimum element
func (q *PriorityQueue) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
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
