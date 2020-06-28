package main

import (
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	// ans := solve(n)
	// ans := solve2(n)
	ans := solve3(n)
	io.Println(ans)
}

// O(n^2 log n) -> TLE
func solve(n int) (ans int) {
	for i := 1; i <= n; i++ {
		count := countDivisors(i)
		ans += i * count
	}
	return ans
}

func countDivisors(n int) int {
	primeCounts := map[int]int{}
	for i := 2; i <= n; i++ {
		count := 0
		for n%i == 0 {
			n /= i
			count++
		}
		if count > 0 {
			primeCounts[i] = count
		}
	}
	ret := 1
	for _, count := range primeCounts {
		ret *= count + 1
	}
	return ret
}

// O(n log n)
func solve2(n int) (ans int) {
	counts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			counts[j]++
		}
	}
	for i, v := range counts {
		ans += i * v
	}
	return ans
}

// O(n)
func solve3(n int) (ans int) {
	for i := 1; i <= n; i++ {
		ans += i * csum(n/i)
	}
	return ans
}

func csum(n int) int {
	return n * (n + 1) / 2
}
