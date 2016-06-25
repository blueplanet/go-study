package main

import (
  "fmt"
)

func test(x, y int) int {
  if x2 := x * x; x2 < y {
    return x2
  } else {
    fmt.Println("x2 < y")
    fmt.Println(x2)
  }

  // if-else.go:15: undefined: x2
  // fmt.Println(x2)

  return y
}

func main() {
  fmt.Println(
    test(2, 5),
    test(3, 6),
  )
}
