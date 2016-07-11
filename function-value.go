package main

import (
	"fmt"
)

func compute(fn func(float64, float64) float64, x, y float64) float64 {
	return fn(x, y)
}

func main() {
	hypot := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot, 5, 12))
}
