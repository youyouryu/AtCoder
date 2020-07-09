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
	line := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(line[0])
	a, _ := strconv.Atoi(line[1])
	b, _ := strconv.Atoi(line[2])
	ans := solve(n, a, b)
	fmt.Println(ans)
}

func solve(n, a, b int) (ans int) {
	// O(n log n)
	for i := 1; i <= n; i++ {
		s := sumOfDigits(i)
		if a <= s && s <= b {
			ans += i
		}
	}
	return ans
}

func sumOfDigits(n int) (ans int) {
	// O(log n)
	for {
		digit := n % 10
		ans += digit
		n /= 10
		if n == 0 {
			break
		}
	}
	return ans
}
