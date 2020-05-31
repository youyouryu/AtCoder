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
	h1, _ := strconv.Atoi(line[0])
	m1, _ := strconv.Atoi(line[1])
	h2, _ := strconv.Atoi(line[2])
	m2, _ := strconv.Atoi(line[3])
	k, _ := strconv.Atoi(line[4])
	t1 := toMinutes(h1, m1)
	t2 := toMinutes(h2, m2)
	fmt.Println(t2 - t1 - k)
}

func toMinutes(h, m int) int {
	return h*60 + m
}
