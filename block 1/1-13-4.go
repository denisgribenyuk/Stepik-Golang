package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var c int
	fmt.Scan(&c)
	is_tr := a*a+b*b == c*c
	if is_tr {

		fmt.Println("Прямоугольный")
	} else {

		fmt.Println("Непрямоугольный")
	}
}
