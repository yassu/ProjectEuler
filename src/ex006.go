package main

import (
	"fmt"
)

func main() {
	N := 100

	// 和の計算
	s1 := 0
	for j := 1; j <= N; j++ {
		s1 += j
	}

	// 2乗の和の計算
	s2 := 0
	for j := 1; j <= N; j++ {
		s2 += j * j
	}

	fmt.Println(s1*s1 - s2)
}
