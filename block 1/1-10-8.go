package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var x int
	var y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	var y_to_str = strconv.Itoa(y)
	var x_to_str = strconv.Itoa(x)
	var result strings.Builder
	for _, symbol := range x_to_str {
		if strings.Contains(y_to_str, string(symbol)) {
			result.WriteString(string(symbol) + " ")
		}
	}
	fmt.Println(result.String())
}





