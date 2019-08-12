package main

import (
    "fmt"
    "errors"
    "math"
)

func sqrt(value float64) (float64,error) {
    //fmt.Println(value)
    if value < 0 {
        return 0, errors.New("Math: -ve number passed")
    }
    return math.Sqrt(value),nil
}

func main() {
    result , err := sqrt(-64)
    if err != nil {
        fmt.Println(err)
    }else {
        fmt.Println("square root :",result,err)
    }

    result , err = sqrt(64)
    if err != nil {
        fmt.Println(err)
    }else {
        fmt.Println("square root :",result,err)
    }
}