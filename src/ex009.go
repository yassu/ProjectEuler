package main

import (
	"fmt"
)

func main() {
	MaxN := 1000

	var c int
	for a := 1; a < MaxN; a++ {
		for b := a + 1; b < MaxN; b++ {
			if a+b >= MaxN {
				break
			}
			c = MaxN - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				break
			}
		}
	}
}
