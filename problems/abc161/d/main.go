package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	k := io.NextInt()
	q := Queue{}
	for i := 1; i <= 9; i++ {
		q.Push(i)
	}
	for i := 0; i < k-1; i++ {
		x := q.Pop()
		if x%10 != 0 {
			q.Push(10*x + x%10 - 1)
		}
		q.Push(10*x + x%10)
		if x%10 != 9 {
			q.Push(10*x + x%10 + 1)
		}
	}
	io.Println(q.Pop())
}

// Queue for Lunlun number
type Queue []int

// Push appends x to end of queue
func (q *Queue) Push(x int) { *q = append(*q, x) }

// Pop returns first item
func (q *Queue) Pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
