package main

import "fmt"

func main() {
  m := make(map[string]string)
  m["name"] = "gyo"

  fmt.Println(m["name"])
}
