package main

import "fmt"
func main(){

  var wordVariants [3]string = [3]string{"korova", "korovy", "korov"}
  var a int
  fmt.Scan(&a)
  lastNum := a % 10
  firstNum := a / 10
  if lastNum ==  1 && firstNum  !=1{
fmt.Printf("%d %s", a,wordVariants[0])
} else if (lastNum == 2||lastNum== 3|| lastNum==4) && firstNum != 1{
fmt.Printf("%d %s",a,wordVariants[1])
} else  {
fmt.Printf("%d %s",a,wordVariants[2])
}
}