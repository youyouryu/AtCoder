package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	t, p := []int{}, []point{}
	for i := 0; i < n; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		ti, _ := strconv.Atoi(line[0])
		xi, _ := strconv.Atoi(line[1])
		yi, _ := strconv.Atoi(line[2])
		t = append(t, ti)
		pi := point{x: xi, y: yi}
		p = append(p, pi)
	}
	ans := solve(t, p)
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func solve(t []int, p []point) bool {
	from := point{x: 0, y: 0}
	now := 0
	for i := range t {
		if !canVisitInTime(from, p[i], t[i]-now) {
			return false
		}
		from = p[i]
		now = t[i]
	}
	return true
}

type point struct {
	x int
	y int
}

func manhattan(p1, p2 point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

func canVisitInTime(from, to point, time int) bool {
	dist := manhattan(from, to)
	rest := time - dist
	if rest < 0 {
		return false
	}
	return rest%2 == 0
}
