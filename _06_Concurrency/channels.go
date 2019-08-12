package main

import (
    "fmt"
)

func sum(s []int, c chan int) {
    sum := 0
    for _ , v := range s {
        sum = sum + v
    }
    fmt.Println("data is ",sum)
    c <- sum                    //write(send values) value to channel and main gorouting block until some goroutine read it
}

func main() {               //active goroutine default main
    s := []int{7,2,8,-9,4,0}
    c := make(chan int, 3)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)

    x, y := <-c , <-c

    fmt.Println(x,y,x+y)
}