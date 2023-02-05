package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	var res int
	if a%9 == 0 {
		res = 9
	} else {
		res = a % 9
	}
	if a < 9 {
		res = a
	}
	fmt.Println(res)
}





