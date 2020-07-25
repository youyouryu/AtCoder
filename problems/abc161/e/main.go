package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n, k, c := io.NextInt(), io.NextInt(), io.NextInt()
	s := io.Next()

	fw := greedy(n, k, c, s)
	bw := greedy(n, k, c, reverse(s))
	for i := 0; i < k/2; i++ {
		bw[i], bw[k-1-i] = bw[k-1-i], bw[i]
	}
	for i := 0; i < k; i++ {
		bw[i] = n - 1 - bw[i]
	}
	for i := range fw {
		if fw[i] == bw[i] {
			io.Println(fw[i] + 1)
		}
	}
}

func greedy(n, k, c int, s string) []int {
	ret := make([]int, k)
	for i := range ret {
		ret[i] = -1
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == 'x' {
			continue
		}
		ret[cnt] = i
		cnt++
		if cnt == k {
			break
		}
		i += c
	}
	return ret
}

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
