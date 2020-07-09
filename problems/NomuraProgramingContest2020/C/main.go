package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	d, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	line := strings.Split(sc.Text(), " ")
	a := []int{}
	for i := range line {
		ai, _ := strconv.Atoi(line[i])
		a = append(a, ai)
	}
	fmt.Println(d, a)
	m := maxNumOfVertices(d, a)
	fmt.Println(m)
}

func maxNumOfVertices(depth int, numOfLeaves []int) (m int) {
	n := len(numOfLeaves)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		dp[i+1] = (dp[i] - numOfLeaves[i]) * 2
		fmt.Println(dp)
		if dp[i+1] < 0 {
			return -1
		}
	}
	for i := 0; i < n-1; i++ {
		m += dp[i]
	}
	m += numOfLeaves[n-1]
	return m
}
