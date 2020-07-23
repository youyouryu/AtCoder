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
	n, k := io.NextInt(), io.NextInt()
	cnt := make([]int, k+1)
	for x := k; x >= 1; x-- {
		cnt[x] = ModPow(k/x, n, mod)
		for y := 2 * x; y <= k; y += x {
			cnt[x] += mod - cnt[y]
			cnt[x] %= mod
		}
	}
	ans := 0
	for x := 1; x <= k; x++ {
		ans += x * cnt[x]
		ans %= mod
	}
	io.Println(ans)
}

const mod = 1e9 + 7

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

// ModMul returns a*b with mod
func ModMul(a, b, mod int) int {
	a, b = a%mod, b%mod
	if b == 0 {
		return 0
	}
	if a*b/b == a {
		return a * b % mod
	}
	panic("overflow")
}

// ModPow returns base^n with mod
func ModPow(base, n, mod int) int {
	ans := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ans = ModMul(ans, base, mod)
		}
		base = ModMul(base, base, mod)
	}
	return ans
}
