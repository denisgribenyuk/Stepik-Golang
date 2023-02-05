package main

import "fmt"

func main() {
	countFib := 2
	var inpNum int
	fmt.Scan(&inpNum)
	if inpNum == 0 {
		fmt.Println(0)
	} else if inpNum == 1 {
		fmt.Println(1)
	} else {
		fib := 2
		prevFib := 1
		prevPrevFib := 1
		for {
			fib = prevPrevFib + prevFib
			prevPrevFib = prevFib
			prevFib = fib
			countFib++
			if fib > inpNum {
				fmt.Println(-1)
				break
			} else if fib == inpNum {
				fmt.Println(countFib)
				break
			}
		}
	}

}





