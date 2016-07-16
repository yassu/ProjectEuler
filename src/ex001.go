package main

import (
	"fmt"
)

func main() {
	END_NUM := 1000
	DIV1 := 3
	DIV2 := 5

	s := 0
	for j := 1; j < END_NUM; j++ {
		if j%DIV1 == 0 || j%DIV2 == 0 {
			s += j
		}
	}
	fmt.Println(s)
}
