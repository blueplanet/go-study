package main

import "fmt"

type Vertex struct {
  Name string
  Age int
}

func main() {
  v := Vertex{"gyo", 36}
  v.Name = "Yu"
  v.Age = 35
  fmt.Println(v)
}
