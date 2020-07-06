package main

import (
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
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
