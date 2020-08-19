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
	R, C, K := io.NextInt(), io.NextInt(), io.NextInt()
	values := make([][]int, R)
	for i := range values {
		values[i] = make([]int, C)
	}
	for i := 0; i < K; i++ {
		ri, ci, vi := io.NextInt()-1, io.NextInt()-1, io.NextInt()
		values[ri][ci] = vi
	}
	ans := solve(values)
	io.Println(ans)
}

func solve(values [][]int) int {
	R, C, L := len(values), len(values[0]), 3
	dp := make([][][]int, R)
	for i := range dp {
		dp[i] = make([][]int, C)
		for j := range dp[i] {
			dp[i][j] = make([]int, L+1)
		}
	}
	dp[0][0][1] = values[0][0]
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			for l := 0; l < L+1; l++ {
				if i+1 < R {
					dp[i+1][j][0] = Max(dp[i+1][j][0], dp[i][j][l])
					dp[i+1][j][1] = Max(dp[i+1][j][1], dp[i][j][l]+values[i+1][j])
				}
				if j+1 < C {
					dp[i][j+1][l] = Max(dp[i][j+1][l], dp[i][j][l])
					if l+1 < L+1 {
						dp[i][j+1][l+1] = Max(dp[i][j+1][l+1], dp[i][j][l]+values[i][j+1])
					}
				}
			}
		}
	}
	return Max(dp[R-1][C-1]...)
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
