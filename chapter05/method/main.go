package main

import (
  "fmt"
)

/* 構造体 Point */
type Point struct{ X, Y int }

/* *Point型のメソッドRender */
func (p *Point) Render() {
  fmt.Printf("<%d, %d>", p.X, p.Y)
}

func main() {

  p := &Point{X: 5, Y: 12}

  p.Render()

}
