package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	for i := 1; i < 180; i++ {
		ans := solve(i)
		fmt.Println(i, ans)
	}
}
