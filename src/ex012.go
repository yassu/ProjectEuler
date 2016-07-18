package main

import (
	"fmt"
)

func factor(n int) map[int]int {
	m := map[int]int{}
	for {
		for j := 2; j <= n; j++ {
			if n%j == 0 {
				n /= j
				m[j] += 1
				break
			}
		}
		if n == 1 {
			return m
		}
	}
}

func numberOfDivisables(n int) int {
	factors := factor(n)
	prod := 1
	for _, pow := range factors {
		prod *= pow + 1
	}
	return prod
}

func main() {
	var numberOf int
	var tri int

	n := 1
	for {
		tri = n * (n + 1) / 2
		numberOf = numberOfDivisables(tri)
		if numberOf >= 500 {
			fmt.Println(numberOf)
			fmt.Println(tri)
			break
		}
		n += 1
	}
}
