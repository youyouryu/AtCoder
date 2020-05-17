package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	switch n % 10 {
	case 2, 4, 5, 7, 9:
		fmt.Println("hon")
	case 0, 1, 6, 8:
		fmt.Println("pon")
	case 3:
		fmt.Println("bon")
	}
}
