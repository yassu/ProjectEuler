package main

// 実際には間違い
import (
	"fmt"
)

func collatz_len(n int) int {
	var l int
	for {
		if n == 1 {
			return l + 1
		}

		l += 1
		if n%2 == 0 {
			n /= 2
		} else if n%2 == 1 {
			n = 3*n + 1
		}
	}
}

func main() {
	maxN := 1000000

	var maxCol int
	for j := 1; j < maxN; j++ {
		l := collatz_len(j)
		if l > maxCol {
			fmt.Println(j)
			maxCol = j
		}
	}
	fmt.Println(maxCol)
}
