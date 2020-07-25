package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	x := io.NextInt()
	ans := 0
	ans += x / 500 * 1000
	x %= 500
	ans += x / 5 * 5
	x %= 5
	io.Println(ans)
}
