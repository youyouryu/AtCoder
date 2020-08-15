package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	k := io.NextInt()
	pow := 1
	num := 7
	history := map[int]struct{}{}
	for {
		if _, ok := history[num]; ok {
			io.Println(-1)
			return
		}
		history[num] = struct{}{}
		if num%k == 0 {
			break
		}
		pow = 10 * pow % k
		num += 7 * pow % k
		num %= k
	}
	io.Println(len(history))
}
