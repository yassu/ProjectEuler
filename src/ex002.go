package main

import (
	"fmt"
)

func main() {
	end_num := 4 * 1000000

	s := 0
	f1, f2 := 1, 1
	for f1 <= end_num {
		if f1%2 == 0 {
			s += f1
		}
		f1, f2 = f2, f1+f2
	}
	fmt.Println(s)
}
