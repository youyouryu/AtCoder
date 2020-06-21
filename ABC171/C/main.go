package main

import (
	"fmt"
	"os"

	"github.com/youyouryu/AtCoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	ans := solve(n)
	fmt.Println(ans)
}

func solve(n int) (ans string) {
	divisor := 26
	rs := []int{}
	for {
		n--
		r := n % divisor
		n /= divisor
		rs = append(rs, r)
		if n == 0 {
			break
		}
	}
	return toString(rs)
}

func toRune(r int) rune {
	return 'a' + rune(r)
}

func toString(rs []int) string {
	runes := make([]rune, len(rs))
	for i := range runes {
		runes[i] = toRune(rs[len(rs)-1-i])
	}
	return string(runes)
}
