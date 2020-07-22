package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	io := NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, m, s := io.NextInt(), io.NextInt(), io.NextInt()
	u, v, a, b := make([]int, m), make([]int, m), make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		u[i], v[i], a[i], b[i] = io.NextInt()-1, io.NextInt()-1, io.NextInt(), io.NextInt()
	}
	c, d := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		c[i], d[i] = io.NextInt(), io.NextInt()
	}
	ans := solve(s, u, v, a, b, c, d)
	for i := range ans {
		io.Println(ans[i])
	}
}

func solve(s int, u, v, a, b, c, d []int) []int {
	n, m := len(c), len(u)
	move := make([]map[int]Cost, n)
	for i := range move {
		move[i] = map[int]Cost{}
	}
	for i := 0; i < m; i++ {
		from, to, cost := u[i], v[i], Cost{Time: b[i], Silver: a[i]}
		move[from][to] = cost
		move[to][from] = cost
	}
	exchange := make([]Cost, n)
	for i := 0; i < n; i++ {
		cost := Cost{Time: d[i], Silver: c[i]}
		exchange[i] = cost
	}

	maxSilver := Max(a...) * (n - 1)
	if s > maxSilver {
		s = maxSilver
	}
	minTime := make([][]int, n)
	for i := range minTime {
		minTime[i] = make([]int, maxSilver+1)
		for j := range minTime[i] {
			minTime[i][j] = inf
		}
	}

	q := &PriorityQueue{}

	minTime[0][s] = 0
	next := State{City: 0, Time: 0, Silver: s}
	heap.Push(q, next)
	for q.Len() > 0 {
		curr := heap.Pop(q).(State)
		//logger.Printf("curr = %d %d %d\n", curr.City, curr.Silver, curr.Time)
		// move
		for city, cost := range move[curr.City] {
			time, silver := curr.Time+cost.Time, curr.Silver-cost.Silver
			if silver < 0 {
				continue
			}
			if time >= minTime[city][silver] {
				continue
			}
			minTime[city][silver] = time
			next = State{City: city, Time: time, Silver: silver}
			heap.Push(q, next)
			//logger.Printf("move > %d %d %d\n", next.City, next.Silver, next.Time)
		}
		// exchange
		city, time, silver := curr.City, curr.Time, curr.Silver
		cost := exchange[city]
		for silver+cost.Silver < maxSilver {
			time += cost.Time
			silver += cost.Silver
			if time >= minTime[city][silver] {
				continue
			}
			minTime[city][silver] = time
			next = State{City: city, Time: time, Silver: silver}
			heap.Push(q, next)
			//logger.Printf("exch > %d %d %d\n", next.City, next.Silver, next.Time)
		}
		if silver < maxSilver {
			time, silver = time+cost.Time, maxSilver
			if time >= minTime[city][silver] {
				continue
			}
			minTime[city][silver] = time
			next = State{City: city, Time: time, Silver: silver}
			heap.Push(q, next)
			//logger.Printf("exch > %d %d %d\n", next.City, next.Silver, next.Time)
		}
	}

	ans := make([]int, n)
	for i := range minTime {
		ans[i] = Min(minTime[i]...)
	}
	return ans[1:]
}

const inf = math.MaxInt64

// Cost represents cost for moving or exchanging
type Cost struct{ Time, Silver int }

// State represents state for dijkstra algorithm
type State struct{ City, Time, Silver int }

// Less defines the order to sort
func (s State) Less(t State) bool {
	if s.Time == t.Time {
		return s.Silver >= t.Silver
	}
	return s.Time < t.Time
}

// PriorityQueue is used for dijkstra algorithm

type PriorityQueue []State

func (q PriorityQueue) Len() int           { return len(q) }
func (q PriorityQueue) Less(i, j int) bool { return q[i].Less(q[j]) }
func (q PriorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

// Push adds an element to queue
func (q *PriorityQueue) Push(x interface{}) { *q = append(*q, x.(State)) }

// Pop gets and remove the mimimum element from queue
func (q *PriorityQueue) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}

// Max returns max value of inputs
func Max(a ...int) (max int) {
	for i, ai := range a {
		if i == 0 || ai > max {
			max = ai
		}
	}
	return max
}

// Min returns minimum value of inputs
func Min(a ...int) (min int) {
	for i, ai := range a {
		if i == 0 || ai < min {
			min = ai
		}
	}
	return min
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
