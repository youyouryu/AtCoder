package main

import (
	"os"
	"sort"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, a, b := io.NextInt(), io.NextInt(), io.NextInt()
	p := io.NextInts(n)
	sort.Ints(p)
	i := sort.Search(n, func(i int) bool { return p[i] > a })
	j := sort.Search(n, func(j int) bool { return p[j] > b })
	first := i
	second := j - i
	third := n - j
	ans := lib.Min(first, second, third)
	io.Println(ans)
}
