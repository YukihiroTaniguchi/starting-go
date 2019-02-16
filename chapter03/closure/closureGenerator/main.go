package main

import (
  "fmt"
)

func integers() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {
  i := integers()
  fmt.Printf("%d\n", i())
  fmt.Printf("%d\n", i())
  fmt.Printf("%d\n", i())

  i2 := integers()
  fmt.Printf("%d\n", i2())
}
