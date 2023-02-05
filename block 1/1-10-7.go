package main

import "fmt"

func main() {
	var x int
	var p float32
	var y int
	var counter int
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	for i := x; i < y; i += int(float32(i) * (p / 100)) {
		counter += 1
	}
	fmt.Println(counter)
}





