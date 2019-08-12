package main

import "fmt"


func service() {
    fmt.Println("Service is ruuning")
}
func main() {
    fmt.Println("main started")

    go service()

    select{}                //block the main goroutine

    fmt.Println("main stopped")
}