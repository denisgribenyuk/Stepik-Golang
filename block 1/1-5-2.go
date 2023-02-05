package main

import "fmt"
func main(){

  var a int
  var b int
  fmt.Scan(&a) // считаем переменную 'a' с консоли
  fmt.Scan(&b) // считаем переменную 'b' с консоли

  a *= a
  b *= b
  var c = a + b
  fmt.Println(c)
}