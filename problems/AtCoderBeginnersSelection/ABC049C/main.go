package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Scan()
	s := sc.Text()
	ans := solve(s)
	// ans := solve(s)
	if ans {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(s string) (ans bool) {
	r := regexp.MustCompile(`^(dream|dreamer|erase|eraser)*$`)
	return r.MatchString(s)
}

func solve2(s string) (ans bool) {
	s = strings.Replace(s, "eraser", "_", -1)
	s = strings.Replace(s, "erase", "_", -1)
	s = strings.Replace(s, "dreamer", "_", -1)
	s = strings.Replace(s, "dream", "_", -1)
	s = strings.Replace(s, "_", "", -1)
	return s == ""
}
