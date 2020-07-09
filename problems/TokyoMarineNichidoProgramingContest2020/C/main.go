package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yuyamada/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	k, _ := sc.NextInt()
	a, _ := sc.NextInts(n)
	ans := solve(n, k, a)
	fmt.Println(strings.Join(lib.Ints2Strs(ans), " "))
}

func solve(n, k int, a []int) (ans []int) {
	for i := 0; i < k; i++ {
		a2 := operate(a)
		if sliceEquals(a, a2) {
			break
		}
		a = a2
	}
	return a
}

func operate(a []int) []int {
	n := len(a)
	b := make([]int, n+1)
	for i, ai := range a {
		left, right := lib.Max(0, i-ai), lib.Min(i+ai+1, n)
		b[left]++
		b[right]--
	}
	a2 := make([]int, n)
	a2[0] = b[0]
	for i := 1; i < n; i++ {
		a2[i] = a2[i-1] + b[i]
	}
	return a2
}

func sliceEquals(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
