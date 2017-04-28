package main

import (
  "fmt"
)

func op (operation func (int, int) int, args ...int) int{
  result := args[0]
  for _, value := range args[1:] {
    result = operation(result, value)
  }
  return result
}
func add (a int, b int) int {
  return a + b
}
func sub (a int, b int) int {
  return a - b
}
func mult (a int, b int) int {
  return a * b
}
func main () {
  fmt.Println(op(mult, []int{1,2,3,4}...))
}
