package main

import "fmt"
func main(){

  var a int
  fmt.Scan(&a)

  var hours = a / 30
	var min = (a % 30) * 2
  fmt.Println("It is", hours, "hours", min, "minutes.")
}