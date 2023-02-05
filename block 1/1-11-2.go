package main

import "fmt"

func main() {
	var arrayLen int
	fmt.Scan(&arrayLen)
	someSlice := make([]int, arrayLen, arrayLen)
	for i := 0; i < arrayLen; i++ {
		fmt.Scan(&someSlice[i])
	}
	fmt.Println(someSlice[3])
}
