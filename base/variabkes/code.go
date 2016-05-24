package main
import (
  "fmt"
)

var (
  c int
  d = "world"
  //d string = "world"
)
func main()  {
  var a int
  var bird1,bird2 string
  var bird3,bird4 string = "bird3","bird4" //var bird3,bird4 = "bird3","bird4"
  b := "hello" //var b = "hello" //var b string = "hello"
  _,bird5 := "birdx","bird5"

  fmt.Println(a,b,c,d,bird1,bird2,bird3,bird4,bird5)
}
