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
	canDivide := true
	for canDivide {
		for i := range a {
			if a[i]%2 != 0 {
				canDivide = false
				break
			}
			a[i] /= 2
			if i == len(a)-1 {
				ans++
			}
		}
	}
	return ans
}
