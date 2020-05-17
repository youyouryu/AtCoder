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
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	line := strings.Split(sc.Text(), " ")
	b, _ := strconv.Atoi(line[0])
	c, _ := strconv.Atoi(line[1])
	sc.Scan()
	s := sc.Text()
	fmt.Println(a+b+c, s)
}
