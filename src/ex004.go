package main

import (
	"fmt"
	"strconv"
)

func isKaibun(n int) bool {
	s := ""
	s = strconv.Itoa(n)
	for j := 0; j < len(s); j++ {
		if s[j] != s[len(s)-j-1] {
			return false
		}
	}
	return true
}

func main() {
	maxKaibun := -1
	s := -1

	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			s = i * j
			if isKaibun(s) && s > maxKaibun {
				maxKaibun = s
			}
		}
	}
	fmt.Println(maxKaibun)
}
