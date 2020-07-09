package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	c, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	x, _ := strconv.Atoi(sc.Text())
	ans := solve(a, b, c, x)
	fmt.Println(ans)
}

func solve(a, b, c, x int) (ans int) {
	values := []int{500, 100, 50}
	nums := []int{a, b, c}
	dp := map[int]int{0: 1}
	for i := range values {
		tmp := map[int]int{}
		for j := range dp {
			for k := 1; k <= nums[i]; k++ {
				value := j + k*values[i]
				tmp[value] += dp[j]
			}
		}
		for j := range tmp {
			if j > x {
				continue
			}
			dp[j] += tmp[j]
		}
	}
	return dp[x]
}

func solve2(a, b, c, x int) (ans int) {
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k == x {
					ans++
				}
			}
		}
	}
	return ans
}
