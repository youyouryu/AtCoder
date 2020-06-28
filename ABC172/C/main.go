package main

import (
	"os"
	"sort"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, m, k := io.NextInt(), io.NextInt(), io.NextInt()
	a := io.NextInts(n)
	b := io.NextInts(m)
	// ans := solve(k, a, b)
	ans := solve2(k, a, b)
	io.Println(ans)
}

// O(n log m)
func solve(k int, a, b []int) (ans int) {
	a = append([]int{0}, a...)
	b = append([]int{0}, b...)
	for i := 1; i < len(a); i++ {
		a[i] = a[i-1] + a[i]
	}
	for i := 1; i < len(b); i++ {
		b[i] = b[i-1] + b[i]
	}
	max := 0
	for i := 0; i < len(a); i++ {
		if a[i] > k {
			break
		}
		j := sort.Search(len(b), func(j int) bool { return a[i]+b[j] > k })
		count := i + j - 1
		if count > max {
			max = count
		}
	}
	return max
}

// O(n+m): two pointers (尺取法)
func solve2(k int, a, b []int) (maxCount int) {
	sum, j := 0, 0
	for j < len(b) && sum+b[j] <= k {
		sum += b[j]
		j++
	}
	j--
	maxCount = j + 1
	for i := 0; i < len(a); i++ {
		sum += a[i]
		for j >= 0 && sum > k {
			sum -= b[j]
			j--
		}
		if sum > k {
			break
		}
		count := i + j + 2
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
