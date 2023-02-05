package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var n1 = a / 100000
	var n2 = a / 10000 % 10
	var n3 = a / 1000 % 10
	var n4 = a / 100 % 10
	var n5 = a / 10 % 10
	var n6 = a % 10
	if n1+n2+n3 == n4+n5+n6 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}