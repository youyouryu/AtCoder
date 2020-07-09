package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

func main() {
	sc := newScanner()
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	t := sc.Text()
	ans := solve(s, t)
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func solve(s, t string) bool {
	r := regexp.MustCompile(`^` + s + `[a-z]$`)
	return r.MatchString(t)
}

func newScanner() *bufio.Scanner {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	return sc
}
