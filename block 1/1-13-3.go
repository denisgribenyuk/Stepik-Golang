package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	hour := num / 3600
	minite := num % 3600 / 60
	fmt.Printf("It is %d hours %d minutes.", hour, minite)
}





