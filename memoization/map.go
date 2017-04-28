package main

import (
  "fmt"
)

func main () {
  memo := make(map[string] int)
  memo["key"] = 1
  fmt.Println(memo["key"])
  fmt.Println(memo["unknown_prop"])
}
