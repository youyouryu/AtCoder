package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	n int
	a []int
	s string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	t, _ := strconv.Atoi(sc.Text())
	testCases := []testCase{}
	for i := 0; i < t; i++ {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		a := []int{}
		for j := 0; j < n; j++ {
			aj, _ := strconv.Atoi(line[j])
			a = append(a, aj)
		}
		sc.Scan()
		s := sc.Text()
		testCases = append(testCases, testCase{n, a, s})
	}
	out := ""
	for i := range testCases {
		ans := solve(testCases[i])
		out += fmt.Sprintln(ans)
	}
	fmt.Print(out)
}

func solve(tc testCase) (ans int) {
	n := tc.n
	a := tc.a
	s := []rune(tc.s)
	tbl := table{0: true}
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			tbl.update(a[i])
			continue
		}
		if !tbl.exists(a[i]) {
			return 1
		}
	}
	return 0
}

type table map[int]bool

func (t *table) update(num int) {
	if t.exists(num) {
		return
	}
	for k := range *t {
		(*t)[k^num] = true
	}
}

func (t *table) exists(num int) bool {
	_, ok := (*t)[num]
	return ok
}
