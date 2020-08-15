package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	a := io.NextInts(n)
	x := 1000
	for i := 0; i < n-1; i++ {
		if a[i+1] > a[i] {
			k := x / a[i]
			x %= a[i]
			x += k * a[i+1]
		}
	}
	io.Println(x)
}

func main2() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	n := io.NextInt()
	a := io.NextInts(n)
	b, c := []int{}, []int{}
	for i := 0; i < n-1; i++ {
		for i < n-1 && a[i+1] <= a[i] {
			i++
		}
		min := a[i]
		if i == n-1 {
			break
		}
		for i < n-1 && a[i+1] >= a[i] {
			i++
		}
		max := a[i]
		b = append(b, min)
		c = append(c, max)
	}
	money, stock := 1000, 0
	for i := range b {
		buy := money / b[i]
		money -= buy * b[i]
		stock += buy

		sell := stock
		money += sell * c[i]
		stock -= sell
	}
	io.Println(money)
}
