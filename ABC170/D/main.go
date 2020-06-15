package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	a, _ := sc.NextInts(n)
	fmt.Println(solve(n, a))
}

func solve(n int, a []int) (ans int) {
	max := lib.Max(a...)
	divisorCount := make([]int, max+1)

	for _, ai := range a {
		for j := ai; j <= max; j += ai {
			divisorCount[j]++
		}
	}

	for _, ai := range a {
		if divisorCount[ai] == 1 {
			ans++
		}
	}

	return ans
}

func solve2(n int, a []int) (ans int) {
	divisible := make(map[int]bool, n)
	for _, ai := range a {
		divisible[ai] = false
	}

	sort.Ints(a)
	max := a[n-1]
	for i, ai := range a {
		if divisible[ai] {
			continue
		}
		for j := ai; j <= max; j += ai {
			if _, ok := divisible[j]; ok {
				divisible[j] = true
			}
		}
		if i < n-1 && a[i+1] == ai {
			continue
		}
		ans++
	}
	return ans
}
