package main

import "fmt"
func main(){

  var a int
  fmt.Scan(&a)
  firstNum := 1
  for {
    if firstNum * 2 > a {
      break  
    }
    fmt.Printf("%d ",firstNum)
    firstNum = firstNum * 2
  }
    fmt.Printf("%d ",firstNum)
}





