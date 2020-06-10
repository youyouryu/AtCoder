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
	k, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()
	fmt.Println(tripleDots(s, k))
}

func tripleDots(s string, k int) string {
	chars := []rune(s)
	if len(chars) <= k {
		return s
	}
	return string(chars[:k]) + "..."
}
