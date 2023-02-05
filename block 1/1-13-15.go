package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inpNum int
	fmt.Scan(&inpNum)
	var numDel int
	fmt.Scan(&numDel)
	var someArray []int
	ost := 1
	newNum := inpNum
	for {
		newNum = inpNum / ost
		someArray = append([]int{newNum % 10}, someArray...)

		ost *= 10
		if newNum < 10 {
			break
		}
	}
	var filteredArray []string
	for _, el := range someArray {
		if el != numDel {
			filteredArray = append(filteredArray, strconv.Itoa(el))
		}
	}
	res := strings.Join(filteredArray, "")
	result, _ := strconv.Atoi(res)
	fmt.Println(result)
}
