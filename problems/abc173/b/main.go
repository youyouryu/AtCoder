package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	count := make(map[string]int, 4)
	for i := 0; i < n; i++ {
		si := io.Next()
		count[si]++
	}
	keys := []string{"AC", "WA", "TLE", "RE"}
	for _, key := range keys {
		io.Printf("%s x %d\n", key, count[key])
	}
}
