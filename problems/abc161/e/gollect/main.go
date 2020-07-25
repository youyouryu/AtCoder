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
	n, k, c := io.NextInt(), io.NextInt(), io.NextInt()
	s := io.Next()

	fw := greedy(n, k, c, s)
	bw := greedy(n, k, c, reverse(s))
	for i := 0; i < k/2; i++ {
		bw[i], bw[k-1-i] = bw[k-1-i], bw[i]
	}
	for i := 0; i < k; i++ {
		bw[i] = n - 1 - bw[i]
	}
	for i := range fw {
		if fw[i] == bw[i] {
			io.Println(fw[i] + 1)
		}
	}
}

func greedy(n, k, c int, s string) []int {
	ret := make([]int, k)
	for i := range ret {
		ret[i] = -1
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == 'x' {
			continue
		}
		ret[cnt] = i
		cnt++
		if cnt == k {
			break
		}
		i += c
	}
	return ret
}

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
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
