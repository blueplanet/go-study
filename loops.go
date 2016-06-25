package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := 1.0
  last := z
  limit := x

  for limit > 0.00000000000000001 {
    z = z - ( z * z - x) / ( 2 * z)

    limit = math.Abs(last - z)
    last = z
  }

  return z
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(math.Sqrt(2))
}
