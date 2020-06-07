package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Split(sc.Text(), " ")
	a, _ := strconv.ParseFloat(line[0], 64)
	b, _ := strconv.ParseFloat(line[1], 64)
	fmt.Println(int(a * b))
}
