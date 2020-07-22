package main

import (
	"math"
	"os"

	"github.com/yuyamada/atcoder/lib"
)

func main() {
	io := lib.NewIo(os.Stdin, os.Stdout)
	defer io.Flush()
	r := io.NextInt()
	io.Println(2 * math.Pi * float64(r))
}
