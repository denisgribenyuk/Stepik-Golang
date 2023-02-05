package main

import "fmt"

func main() {
	var lenArray int
	fmt.Scan(&lenArray)
	var slice = make([]int, lenArray, lenArray)
	var a int
	for i := 0; i < lenArray; i++ {
		fmt.Scan(&a)
		slice[i] = a
	}
	posit_count := 0
	for _, el := range slice {
		if el >= 0 {
			posit_count ++
		}
	}
	fmt.Println(posit_count)
}
