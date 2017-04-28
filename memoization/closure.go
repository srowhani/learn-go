package main;
import (
  "fmt"
  "os"
)

func d(array []int) func() int {

  return func () int{
    l := len(array)
    if l < 1 {
      return -1
    }
    item := array[l - 1]
    array = array[0 : l - 1]
    return item
  }
}
func main () {
  a := d([]int{1,2,3})
  var result int
  for {
    result = a()
    switch result {
      case -1:
        os.Exit(1)
      default:
        fmt.Println(result)
    }
  }

  fmt.Println(a(), a(), a(), a())
}
