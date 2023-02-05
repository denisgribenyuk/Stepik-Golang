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
	for i, el := range slice {
		if i%2 == 0 {
			fmt.Printf("%d ", el)
		}
	}
}
