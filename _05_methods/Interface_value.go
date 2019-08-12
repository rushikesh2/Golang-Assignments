package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64

}

type Rect struct {
    height float64
    width float64
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func (r Rect) Area() float64 {
    return r.height * r.width
}
func main() {
    var s Shape
    s = Circle{10}                  //concrete type of s is Circle and it implements Area()
    fmt.Printf("type of s is %T\n", s)
    fmt.Printf("value of s is %v\n", s)
    fmt.Println("Area of Circle is :", s.Area())

     fmt.Println("**********************************")

    s = Rect{3 , 5}                 //concrete type of s is Rect and it implements Area()
    fmt.Printf("type of s is %T\n", s)
    fmt.Printf("value of s is %v\n", s)
    fmt.Println("Area of Rectangal is :", s.Area())
}