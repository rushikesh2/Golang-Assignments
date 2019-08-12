package main

import "fmt"

var a int = 0
func square(c chan int) {
    for i := 0 ; i < 4 ; i++ {
        num := <- c
        fmt.Println(num * num)
    }
    close(c)
}
func main() {

    fmt.Println("Buffer example")
    c := make(chan int,3)

    go square(c)
    c <- 1
    c <- 2
    c <- 3
    c <- 4
    fmt.Println("Buffer program executed")
}