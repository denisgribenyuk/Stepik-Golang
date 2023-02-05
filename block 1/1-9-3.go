package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)
	var str_len = len(strconv.Itoa(a))
	switch str_len {
	case 1:
		fmt.Println(a)
	case 2:
		fmt.Println(a / 10)
	case 3:
		fmt.Println(a / 100)
	case 4:
		fmt.Println(a / 1000)
	case 5:
		fmt.Println(a / 10000)
	}
}
