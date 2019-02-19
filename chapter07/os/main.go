package main

import (
  "fmt"
  "log"
  "os"
)

func main() {
  defer func() {
    fmt.Println("defer")
  }()

  _, err := os.Open("foo")
  if err != nil {
    log.Fatal(err)
  }

  os.Exit(0)

}
