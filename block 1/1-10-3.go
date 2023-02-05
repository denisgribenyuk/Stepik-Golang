package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	var result = 0
	for i := 1; i <= count; i++ {
		var inp int
		fmt.Scan(&inp)
		if inp/100 == 0 && inp/10 != 0 && inp%8 == 0 {
			result += inp
		}
	}
	fmt.Println(result)
}




