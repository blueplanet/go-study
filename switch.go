package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    fmt.Printf("%s.", os)
  }
}