package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Printf("%d", num%10)

	fmt.Printf("%d", num/10%10)
	fmt.Printf("%d", num/100)
}
