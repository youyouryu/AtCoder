package main

import "fmt"

func main() {
	allocateAndPrint(1e9)
	allocateAndPrint(1e10)
	allocateAndPrint(1e11)
	allocateAndPrint(1e12)
	allocateAndPrint(1e13)
	allocateAndPrint(1e14)
	allocateAndPrint(1e15)
}

func allocateAndPrint(size int) {
	a := make([]int, size)
	fmt.Println(len(a))

}
