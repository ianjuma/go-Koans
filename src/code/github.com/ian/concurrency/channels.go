package main

import (
    "fmt"
    )

/*
func ExpensiveComputation(c chan float64) {
    x := 2.345 / 543.3456
    c <- x
}
*/

func GetChanell() {
    channel := make(chan float64)

    go func(c chan float64) {
        x := 20.0/ 2.0
        c <- x
    }(channel)

    value := <-channel
    fmt.Println("Result was:", value)
}

func main() {
    //cha := make(chan float64)
    GetChanell()
}
