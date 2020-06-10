package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)
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
	fmt.Printf(ans)
}

var max = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

func solve(a []int) (ans string) {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sum := big.NewInt(1)
	for i := range a {
		ai := big.NewInt(int64(a[i]))
		sum.Mul(sum, ai)
		if sum.Cmp(max) == 1 {
			return "-1"
		}
	}
	return sum.String()
}
