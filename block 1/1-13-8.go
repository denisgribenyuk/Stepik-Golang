package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	min := 0
	count := 1
	for i := 0; i < a; i++ {

		var b int
		fmt.Scan(&b)

		if i == 0 {

			min = b
		} else {
			if b < min {
				min = b
				count = 1
			} else if b == min {
				count++
			}
		}
	}
	fmt.Println(count)
}
