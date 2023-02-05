package main

import "fmt"
func main(){

  var a int
  fmt.Scan(&a)
  var b int
  fmt.Scan(&b)
  var done bool = false
  for i:=b; i>=a; i-- {
    if i %7 == 0 {
      fmt.Println(i)
      done = !done
      break
    }
  }
  if !done{
  fmt.Println("NO")
}
}