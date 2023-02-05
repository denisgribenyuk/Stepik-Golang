package main

import "fmt"

func main() {
	var count = 1
	var max_num = 0
	var result int
	for count != 0 {
		fmt.Scan(&count)
		if count == max_num {
			result += 1
		} else if count > max_num {
			result = 1
			max_num = count
		}
	}
	fmt.Println(result)
}





