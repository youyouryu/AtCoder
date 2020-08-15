package main

import (
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	x := io.NextInt()
	var grade int
	switch {
	case x < 600:
		grade = 8
	case x < 800:
		grade = 7
	case x < 1000:
		grade = 6
	case x < 1200:
		grade = 5
	case x < 1400:
		grade = 4
	case x < 1600:
		grade = 3
	case x < 1800:
		grade = 2
	case x < 2000:
		grade = 1
	}
	io.Println(grade)
}
