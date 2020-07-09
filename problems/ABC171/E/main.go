package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	n, _ := sc.NextInt()
	a, err := sc.NextInts(n)
	if err != nil {
		panic(err)
	}
	ans := solve(n, a)
	out := make([]interface{}, len(ans))
	for i := range out {
		out[i] = ans[i]
	}
	fmt.Println(out...)
}

func solve(n int, a []int) []int {
	ans := make([]int, n)
	xor := 0
	for i := 0; i < n; i++ {
		xor ^= a[i]
	}
	for i := 0; i < n; i++ {
		ans[i] = xor ^ a[i]
	}
	return ans
}
