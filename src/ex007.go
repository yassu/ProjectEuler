package main

import (
	"fmt"
)

func main() {
	INDEX := 10001

	primes := []int{}

	n := 2
	for len(primes) < INDEX {
		isPrime := true
		for _, prime := range primes {
			if n%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, n)
		}
		n += 1
	}

	fmt.Println(primes[INDEX-1])
}
