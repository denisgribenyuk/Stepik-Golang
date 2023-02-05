package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	c := (float32(a) + float32(b)) / 2
	fmt.Println(c)
}
