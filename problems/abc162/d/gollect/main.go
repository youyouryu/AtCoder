package main

import (
	"bufio"
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
	_ = io.NextInt()
	s := io.Next()
	r, g, b := []int{}, []int{}, []int{}
	for i := range s {
		switch s[i] {
		case 'R':
			r = append(r, i)
		case 'G':
			g = append(g, i)
		case 'B':
			b = append(b, i)
		}
	}
	//logger.Println(r, g, b)
	ans := len(r) * len(g) * len(b)
	ans -= find(r, g, b) + find(r, b, g) + find(g, r, b) + find(g, b, r) + find(b, r, g) + find(b, g, r)
	io.Println(ans)
}

func find(a, b, c []int) (found int) {
	for i := 0; i < len(a); i++ {
		jBegin := sort.Search(len(b), func(j int) bool { return b[j] > a[i] })
		for j := jBegin; j < len(b); j++ {
			kBegin := sort.Search(len(c), func(k int) bool { return c[k] > b[j] })
			k := sort.Search(len(c[kBegin:]), func(k int) bool { return a[i]-2*b[j]+c[kBegin:][k] >= 0 })
			if k < len(c[kBegin:]) && a[i]-2*b[j]+c[kBegin:][k] == 0 {
				found++
			}
		}
	}
	return found
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
