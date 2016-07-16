package main

import (
	"fmt"
)

func getLargestFactor(n int) int {
	f := 1
	for n != 1 {
		j := 2
		for ; j <= n; j++ {
			if n%j == 0 {
				n /= j
				f = j
				break
			}
		}
	}
	return f
}

func main() {
	fmt.Println([]int{1, 2, 3})
	fmt.Println(getLargestFactor(600851475143))
}
