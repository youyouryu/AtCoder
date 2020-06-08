package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(line[0])
	y, _ := strconv.Atoi(line[1])
	i, j, k := solve2(n, y)
	fmt.Println(i, j, k)
}

// O(n^2)
func solve(n, y int) (i, j, k int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k := n - i - j
			if 10000*i+5000*j+1000*k == y {
				return i, j, k
			}
		}
	}
	return -1, -1, -1
}

// O(n log n)
func solve2(n, y int) (i, j, k int) {
	for i := 0; i <= n; i++ {
		j := sort.Search(n-i+1, func(j int) bool {
			k := n - i - j
			return 10000*i+5000*j+1000*k >= y
		})
		if j == n-i+1 {
			continue
		}
		k = n - i - j
		if 10000*i+5000*j+1000*k == y {
			return i, j, k
		}
	}
	return -1, -1, -1
}
