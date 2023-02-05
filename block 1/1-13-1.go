package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	sum := num/100 + num/10%10 + num%10
	fmt.Println(sum)
}
