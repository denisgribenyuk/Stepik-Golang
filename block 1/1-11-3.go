package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	maxNum := array[0]
	for _, el := range array {
		if el > maxNum {
			maxNum = el
		}
	}
	fmt.Println(maxNum)
}
