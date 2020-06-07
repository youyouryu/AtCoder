package main

import (
	"bufio"
	"fmt"
	"math/big"
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
	fmt.Println(multiply(a))
}

func multiply(a []int) (ans string) {
	max := big.NewInt(0)
	max.SetString("1000000000000000000", 10)
	sum := big.NewInt(0)
	for i := 0; i < len(a); i++ {
		ai := big.NewInt(int64(a[i]))
		for j := 0; j < a[i]; j++ {
			sum.Add(sum, ai)
			fmt.Println(sum)
			if max.Cmp(sum) == -1 {
				return "-1"
			}
		}
	}
	return sum.String()
}
