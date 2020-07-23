package main

import (
	"log"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

var logger = log.New(os.Stderr, "", 0)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	_ = io.NextInt()
	s := io.Next()
	ans := solve(s)
	io.Println(ans)
}

func solve(s string) (ans int) {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}
	n := len(s)
	ans = cnt['R'] * cnt['G'] * cnt['B']
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			k := 2*j - i
			if k >= n {
				continue
			}
			if s[i] == s[j] || s[j] == s[k] || s[k] == s[i] {
				continue
			}
			ans--
		}
	}
	return ans
}
