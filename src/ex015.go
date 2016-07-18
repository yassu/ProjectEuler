package main

import (
	"fmt"
)

var M map[[2]int]int

func init() {
	M = map[[2]int]int{}
}

func a(x, y int) int {
	if M[[2]int{x, y}] != 0 {
		return M[[2]int{x, y}]
	}

	if x == 0 || y == 0 {
		M[[2]int{x, y}] = 1
		return 1
	} else {
		M[[2]int{x, y}] = a(x-1, y) + a(x, y-1)
		if x >= 10 && y >= 10 {
		}
		return M[[2]int{x, y}]
	}
}

func main() {
	fmt.Println(a(20, 20))
}
