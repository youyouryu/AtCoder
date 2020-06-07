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
	a, _ := strconv.Atoi(line[0])
	b, _ := strconv.Atoi(line[1])
	fmt.Println(a * b)
}
