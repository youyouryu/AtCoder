package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	sc := lib.NewScanner(os.Stdin)
	n, _ := sc.NextInt()
	k, _ := sc.NextInt()
	p, _ := sc.NextInts(n)
	sort.Ints(p)
	sum := 0
	for i := 0; i < k; i++ {
		sum += p[i]
	}
	fmt.Println(sum)
}
