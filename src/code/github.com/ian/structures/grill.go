// grill game - ping pong
// using goroutines to simulate game events

package main

import (
    "fmt"
    "time"
)

func Ticker() {
    go func() {
        fmt.Println("Goroutine -> Ticker")
    }()
}

func Customer() {
    go func() {
        some := func(x int, y int) (int) {
            return x+y
        }
        fmt.Println(some(10, 23))
    }()
}

func main() {
    fmt.Println("Start -> Main")
    Customer()
    Ticker()
    time.Sleep(2000)
}
