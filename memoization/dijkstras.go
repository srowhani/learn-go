package main;

import (
  "fmt"
  //"math"
  "io/ioutil"
  "encoding/json"
)

type Graph struct {
	V []map[string]int
	E []map[string]int
}
type Results struct {
  m_graph Graph
}
var (
  results Results
  m_graph Graph
)
func init () {
  raw, err := ioutil.ReadFile("graph.json")
  if err == nil {
    json.Unmarshal(raw, &m_graph)
    fmt.Println(m_graph)
  }
}

func main () {
  // d := make(map[int]int)
  // for _, v := range m_graph.V {
  //   // d[v["id"]] = int(math.Inf(+1))
  // }
}
