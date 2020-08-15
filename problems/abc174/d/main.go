package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	s := io.Next()
	cnt := map[byte]int{}
	for i := 0; i < n; i++ {
		cnt[s[i]]++
	}
	ans := 0
	for i := 0; i < cnt['R']; i++ {
		if s[i] == 'W' {
			ans++
		}
	}
	io.Println(ans)
}
