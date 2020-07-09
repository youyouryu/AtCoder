package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	count := 0
	for _, si := range []rune(s) {
		if si == '1' {
			count++
		}
	}
	fmt.Println(count)
}
