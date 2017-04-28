package main;
import (
  "fmt"
  "math"
  "io/ioutil"
  "encoding/json"
)

type Stations []map[string]int
var (
  test_index = 0
  num_scenarios = 0
  stations Stations
  memo = make(map[string] int)
)

func init () {
    raw, err := ioutil.ReadFile("stations.json")
    if err == nil {
      json.Unmarshal(raw, &stations)
      fmt.Println(stations)
    }
}

func main () {
  fmt.Println(opt(len(stations) - 1))
}

func opt(index int) int {
  current_station := stations[index]
  var candidate_station_index int

  if index < 0 {
    return 0
  }
  if index == 0 {
    return stations[index]["revenue"]
  }

  if current_station["location"] - stations[0]["location"] < 20 {
    best_so_far := stations[0]
    for i := 1; i < index; i++ {
      best_candidate := stations[index]
      if best_candidate["revenue"] > best_so_far["revenue"] {
        best_so_far = best_candidate
      }
    }
    return best_so_far["revenue"]
  }


  for i := index -1 ; i >= 0; i-- {
    if current_station["location"] - stations[i]["location"] > 19 {
      candidate_station_index = i
      break
    }
  }
  var opt_i int
  var opt_j int

  key := "opt(" + string(candidate_station_index) + ")"
  _, ok := memo[key]
  if ok {
    opt_j = memo[key]
  } else {
    memo[key] = opt(candidate_station_index)
    opt_j = memo[key]
  }

  key = "opt(" + string(index - 1) + ")"
  _, ok = memo[key]
  if ok {
    opt_i = memo[key]
  } else {
    memo[key] = opt(index - 1)
    opt_i = memo[key]
  }


  return int(math.Max(float64(stations[index]["revenue"]) + float64(opt_i), float64(opt_j)))
}
