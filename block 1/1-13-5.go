package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var c int
	fmt.Scan(&c)
	is_tr := a < b+c && b < a+c && c < a+b
	if is_tr {

		fmt.Println("Существует")
	} else {

		fmt.Println("Не существует")
	}
}





