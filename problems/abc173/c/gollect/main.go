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
	h, w, k := io.NextInt(), io.NextInt(), io.NextInt()
	c := make([]string, h)
	for i := 0; i < h; i++ {
		c[i] = io.Next()
	}
	ans := solve(h, w, k, c)
	io.Println(ans)
}

func solve(h, w, k int, c []string) (ans int) {
	for i := 0; i < 1<<h; i++ {
		for j := 0; j < 1<<w; j++ {
			count := 0
			for is := 0; is < h; is++ {
				if i>>is&1 == 1 {
					continue
				}
				for js := 0; js < w; js++ {
					if j>>js&1 == 1 {
						continue
					}
					if c[is][js] == '#' {
						count++
					}
				}
			}
			if count == k {
				ans++
			}
		}
	}
	return ans
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
