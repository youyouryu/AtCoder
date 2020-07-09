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
	n := io.NextInt()
	a, b := make([]int, n-1), make([]int, n-1)
	for i := 0; i < n-1; i++ {
		ai, bi := io.NextInt(), io.NextInt()
		if ai > bi {
			ai, bi = bi, ai
		}
		a[i], b[i] = ai, bi
	}
	ans := solve(n, a, b)
	io.Println(ans)
}

// O(n)
// 連結成分数 = 頂点数 - 辺数 (閉路がないため)
func solve(n int, a, b []int) (ans int) {
	v, e := 0, 0
	for i := 1; i <= n; i++ {
		v += i * (n - i + 1)
	}
	for i := 0; i < n-1; i++ {
		e += a[i] * (n - b[i] + 1)
	}
	return v - e
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
