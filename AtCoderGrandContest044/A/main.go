package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	n int
	a int
	b int
	c int
	d int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	t, _ := strconv.Atoi(sc.Text())
	tcs := []testCase{}
	for i := 0; i < t; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		n, _ := strconv.Atoi(line[0])
		a, _ := strconv.Atoi(line[1])
		b, _ := strconv.Atoi(line[2])
		c, _ := strconv.Atoi(line[3])
		d, _ := strconv.Atoi(line[4])
		tc := testCase{n, a, b, c, d}
		tcs = append(tcs, tc)
	}
	for _, tc := range tcs {
		ans := solve(tc)
		fmt.Println(ans)
	}
}

func solve(tc testCase) (ans int) {
	size := calcSize(tc.n)
	dp := make([]int, size+1)
	for i := range dp {
		if i == 0 {
			continue
		}
		dp[i] = math.MaxInt64
	}
	for update(tc, dp) {
		fmt.Println(dp)
	}
	return dp[tc.n]
}

func update(tc testCase, dp []int) bool {
	for i := 0; i < len(dp)-1; i++ {
		var cost int
		if (i+1)%2 == 0 {
			cost = dp[(i+1)/2] + tc.a
			if cost < dp[i+1] {
				dp[i+1] = cost
			}
		}
		if (i+1)%3 == 0 {
			cost = dp[(i+1)/3] + tc.b
			if cost < dp[i+1] {
				dp[i+1] = cost
			}
		}
		if (i+1)%5 == 0 {
			cost = dp[(i+1)/5] + tc.c
			if cost < dp[i+1] {
				dp[i+1] = cost
			}
		}
		cost = dp[i] + tc.d
		if cost < dp[i+1] {
			dp[i+1] = cost
		}
	}
	updatedInBackward := false
	for i := len(dp) - 1; i > 0; i-- {
		cost := dp[i] + tc.d
		if cost < dp[i-1] {
			dp[i-1] = cost
			updatedInBackward = true
		}
	}
	return updatedInBackward
}

func calcSize(n int) int {
	div2, div3, div5 := false, false, false
	for i := n; ; i++ {
		if i%2 == 0 {
			div2 = true
		}
		if i%3 == 0 {
			div3 = true
		}
		if i%5 == 0 {
			div5 = true
		}
		if div2 && div3 && div5 {
			return i
		}
	}
}
