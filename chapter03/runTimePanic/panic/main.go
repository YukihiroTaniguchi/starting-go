package main

import (
  "fmt"
)

func main() {
  defer fmt.Println("Hello, World!")
  panic("runtime error!")
}
