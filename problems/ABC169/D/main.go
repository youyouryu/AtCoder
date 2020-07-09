package main

import (
	"fmt"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	ans := solve(n)
	fmt.Println(ans)
}

func solve(n int) (ans int) {
	for i := 2; i*i <= n; i++ {
		for z := i; n%z == 0; z *= i {
			ans++
			n /= z
		}
		for n%i == 0 {
			n /= i
		}
	}
	if n != 1 {
		ans++
	}
	return ans
}
