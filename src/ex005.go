package main

import (
	"fmt"
)

func isAllDivisable(n, maxN int) bool {
	for j := 1; j <= maxN; j++ {
		if n%j != 0 {
			return false
		}
	}
	return true
}

func main() {
	MaxN := 20

	for j := 1; ; j++ {
		if isAllDivisable(j, MaxN) {
			fmt.Println(j)
			break
		}
	}
}
