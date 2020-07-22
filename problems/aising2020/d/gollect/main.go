package main

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	io := NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	_ = io.NextInt()
	s := io.Next()
	ans := solve2(s)
	for _, v := range ans {
		io.Println(v)
	}
}

// pcが変わるのでdrを再利用できない

func solve2(s string) []int {
	n := len(s)
	x := make([]int, n)
	for i := range s {
		x[i] = int(s[i] - '0')
	}
	pc0 := 0
	for i := range x {
		pc0 += x[i]
	}

	ans := make([]int, n)
	// 桁の値でforを回す。
	// 2n回ループを回すが、if文に引っかかるのはn回のみ。
	// 桁の値が同じものをまとめて処理することで、dr%pcを再利用できる
	for b := 0; b < 2; b++ {
		pc := pc0
		if b == 0 {
			pc++
		} else {
			pc--
		}
		if pc <= 0 {
			continue
		}
		r0 := 0
		for i := 0; i < n; i++ {
			r0 = (r0 * 2) % pc
			r0 += x[i]
		}
		dr := 1
		for i := n - 1; i >= 0; i-- {
			r := r0
			if x[i] == b {
				if b == 0 {
					r = (r + dr) % pc
				} else {
					r = (r - dr + pc) % pc
				}
				ans[i] = f(r) + 1
			}
			// 0->1を一括で処理
			// 1->0を一括で処理
			dr = (dr * 2) % pc
		}
	}
	return ans
}

func f(x int) int {
	if x == 0 {
		return 0
	}
	return f(x%pcnt(x)) + 1
}

func pcnt(x int) int {
	return bits.OnesCount(uint(x))
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
