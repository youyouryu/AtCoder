package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	ans := solve(n)
	fmt.Println(ans)
}

func solve(n int) (ans int) {
	p := getPrimes(n)
	l := len(p)
	for i := 0; i < l; i++ {
		for num := p[i] * p[i]; num <= n; num *= p[i] {
			p = append(p, num)
		}
	}
	sort.Slice(p, func(i, j int) bool { return p[i] < p[j] })
	fmt.Println(p)
	for i := range p {
		if n%p[i] != 0 {
			continue
		}
		n /= p[i]
		ans++
	}
	return ans
}

func getPrimes(n int) (ans []int) {
	ans = []int{}
	max := int(math.Sqrt(float64(n)))
	isPrime := make([]bool, max)
	for i := 2; i < max; i++ {
		isPrime[i] = true
	}
	for i := 2; i < max; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i + 1; j < max; j++ {
			if j%i == 0 {
				isPrime[j] = false
			}
		}
		ans = append(ans, i)
	}
	return ans
}
