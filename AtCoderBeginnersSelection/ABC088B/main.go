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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	line := strings.Split(sc.Text(), " ")
	a := []int{}
	for i := 0; i < n; i++ {
		ai, _ := strconv.Atoi(line[i])
		a = append(a, ai)
	}
	ans := solve(a)
	fmt.Println(ans)
}

func solve(a []int) (ans int) {
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	x, y := 0, 0
	for i, v := range a {
		if i%2 == 0 {
			x += v
		} else {
			y += v
		}
	}
	return x - y
}
