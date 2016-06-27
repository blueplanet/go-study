package main

import "fmt"

func main() {
  pow := []int{1, 2, 3, 4, 5, 6, 7, 8}

  for i, v := range pow {
    fmt.Printf("%d %d\n", i, v)
  }
}
