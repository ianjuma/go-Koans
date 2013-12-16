package main

import "fmt"

func main() {

    if num := 12; num < 10 {
        fmt.Println("Num < 10")
    } else if (num > 10) {
        fmt.Println("Num > 10")
    } else {
        fmt.Println("Num == 10")
    }

    for i := 0; i < 100; i++ {

        if i % 3 == 0 {
            fmt.Println("Buzz")

            if i % 5 == 0 {
                fmt.Println("BuzzLightYear")
            } else {
                fmt.Println("LightYear")
            }
        }
    }
}
