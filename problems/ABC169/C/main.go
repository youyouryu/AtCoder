package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := newScanner()
	sc.Scan()
	a, _ := strconv.ParseInt(sc.Text(), 10, 64)
	sc.Scan()
	b, _ := strconv.ParseFloat(sc.Text(), 64)
	ans := a * int64(math.Round(100*b)) / 100
	fmt.Println(ans)
}

func newScanner() *bufio.Scanner {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	return sc
}
