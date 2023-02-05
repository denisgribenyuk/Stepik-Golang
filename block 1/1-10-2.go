package main

import "fmt"

func main() {
	var start int
	fmt.Scan(&start)
	var end int
	fmt.Scan(&end)
	var result = 0
	for i := start; i <= end; i++ {
		result += i
	}
	fmt.Println(result)
}




