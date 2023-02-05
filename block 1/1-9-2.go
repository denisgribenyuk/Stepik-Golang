package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var num1 int = a / 100
	var num2 int = a % 100 / 10
	var num3 int = a % 10

	if num1 != num2 && num1 != num3 && num2 != num3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
