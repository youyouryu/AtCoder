package main

import (
	"log"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

const mod = 2019

var logger = log.New(os.Stderr, "", 0)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	s := io.Next()
	ans := solve(s)
	io.Println(ans)
}

func solve(s string) (ans int) {
	n := len(s)
	a := make([]int, n)
	pow := 1
	for i := range a {
		digit := int(s[n-1-i] - '0')
		a[n-1-i] = digit * pow % mod
		pow = pow * 10 % mod
	}
	cs := make([]int, n+1)
	for i := range a {
		cs[i+1] = (cs[i] + a[i]) % mod
	}
	cnt := map[int]int{}
	for i := range cs {
		ans += cnt[cs[i]]
		cnt[cs[i]]++
	}
	logger.Println("s", s)
	logger.Println("a", a)
	logger.Println("cs", cs)
	logger.Println("cnt", cnt)
	return ans
}
