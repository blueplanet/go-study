package main

import "fmt"

type Vertex struct {
  Name string
  Age int
}

func main() {
  fmt.Println(Vertex{"gyo", 36})
}
