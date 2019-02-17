package main

import (
  "fmt"
  "math"
)

type IntPoint struct{ X, Y int }
type FloatPoint struct{ X, Y float64 }

func (p *IntPoint) Distance(dp *IntPoint) float64 {
  x, y := p.X-dp.X, p.Y-dp.Y
  return math.Sqrt(float64(x*x + y*y))
}

func (p *FloatPoint) Distance(dp *FloatPoint) float64 {
  x, y := p.X-dp.X, p.Y-dp.Y
  return math.Sqrt(x*x + y*y)
}

func main() {
  ip := &IntPoint{X: 1, Y: 2}
  fp := &FloatPoint{X: 1.5, Y: 2.5}

  disi := ip.Distance(&IntPoint{X: 3, Y: 5})
  fmt.Println(disi)

  disf := fp.Distance(&FloatPoint{X: 3.5, Y: 5.5})
  fmt.Println(disf)
}
