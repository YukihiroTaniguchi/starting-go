package main

import (
  "fmt"
  "log"
  "time"
)

func main() {
  t := time.Now()
  fmt.Println(t)
  fmt.Printf("%d == %s == %s\n", t.Month(), t.Month(), t.Month().String())

  d, _ := time.ParseDuration("2h30m")
  fmt.Println(d)

  t = t.Add(2*time.Hour + 2*time.Minute)
  fmt.Println(t)

  t0 := time.Now()
  t1 := time.Now()
  fmt.Println(t0.Sub(t1))

  fmt.Println(t1.Before(t0))
  fmt.Println(t1.After(t0))

  t2 := t0.AddDate(0, 2, -12)
  fmt.Println(t2)

  t3, err := time.Parse("2006/01/02", "2015/11/27")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(t3)

  t4 := time.Now()
  unix := t4.Unix()
  fmt.Println(unix)

  t5 := time.Unix(unix, 0)
  fmt.Println(t5)

  ch := time.After(5 * time.Second)
  fmt.Println(<-ch)

}
