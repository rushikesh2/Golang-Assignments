package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X , Y float64
}

/*func abs(v Vertex) float64 {                        //function with argument
    fmt.Println("abs method",v.Y)
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}*/

func (v Vertex) abs() float64 {                        //function wiht argument
    fmt.Println("Modify the values :",v.X)
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {                 //pointer receiver
    fmt.Println("initail value :",v.X)
    v.X = v.X * f
    fmt.Println("values after Mul:" , v.X)
    v.X = v.X + 1
    fmt.Println("value after addtion :", v.X)
    v.Y = v.Y * f
}
func main() {
    v := Vertex{3 , 4}
    Scale(&v , 10)
    v.abs()
}

