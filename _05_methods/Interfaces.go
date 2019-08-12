package main

import (
    "fmt"
)

type Shape interface {
    Area() float64
}

type T struct {
    height float64
    width float64
}

func (t T)  Area() float64 {
    return t.height * t.width
}

func describe(s Shape) {
	fmt.Printf("(%v, %T)\n", s, s)
}

func main() {
    var s Shape
    s = T{4 , 5}
    t := T{4 , 5}         //implement this way also
  //s = t
    describe(s)
    fmt.Println("area of Rectangle  : ",s.Area())
    fmt.Println("s == t is :", s == t)
}