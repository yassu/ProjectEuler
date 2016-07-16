package main

import (
	"fmt"
)

func main() {
	MaxN := 2000000

	primes := []int{}
	n := 2
	for {
		isPrime := true
		for _, prime := range primes {
			if n%prime == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			if n > MaxN {
				break
			}
			primes = append(primes, n)
			fmt.Println(n)
		}
		n += 1
	}

	s := 0
	for _, num := range primes {
		s += num
	}
	fmt.Println(s)
}
