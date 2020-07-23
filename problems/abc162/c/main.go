package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	k := io.NextInt()
	ans := 0
	for a := 1; a <= k; a++ {
		for b := 1; b <= k; b++ {
			gcd := lib.Gcd(a, b)
			for c := 1; c <= k; c++ {
				ans += lib.Gcd(gcd, c)
			}
		}
	}
	io.Println(ans)
}
