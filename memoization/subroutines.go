package main;
import (
  "fmt"
  "time"
)
func init () {

}

func main () {
  go func () {
    for i := 0; i < 3; i++ {
      fmt.Println("async ", i)
      time.Sleep(time.Second * 1)
    }
  }()

  time.Sleep(time.Second / 2)
  
  for i:=0;i<3;i++{
    fmt.Println("sync ", i)
    time.Sleep(time.Second * 1)
  }
}
