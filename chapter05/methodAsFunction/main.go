package main

import (
  "fmt"
)

type Point struct{ X, Y int }

func (p *Point) ToString() string {
  return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}

func main() {
  f := (*Point).ToString
  fmt.Println(f(&Point{X: 1, Y: 2}))

  fmt.Println((*Point).ToString(&Point{X: 1, Y: 2}))

  p := &Point{X: 1, Y: 2}
  f2 := p.ToString
  fmt.Println(f2())

}
