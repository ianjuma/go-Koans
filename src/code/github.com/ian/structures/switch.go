package main

import (
    "time"
    "fmt"
)

func main() {
    i := int(0)
    _, err := fmt.Scanf("%i", &i)

    if err == nil {
        panic("ERR")
    }

    switch {
        case i > 10:
            fmt.Println("i > 10")

        case i > 100:
            fmt.Println("i > 100")

        case i < 10:
            fmt.Println("i < 10")

        default:
            fmt.Println("I don't know what that is")

    }

    switch time.Now().Weekday() {
        case time.Saturday, time.Sunday, time.Friday:
            fmt.Println("Root me a Beer!")

        case time.Thursday:
            t := time.Now()

            switch {
                case t.Hour() > 16:
                    fmt.Println("After class ... ")

                case t.Hour() < 15:
                    fmt.Println("You have a class to attend")
            }

        default:
            fmt.Println("Go to class")
    }


    t := time.Now()
    switch {
        case t.Hour() < 12:
            fmt.Println("Morning")

        default:
            fmt.Println("Afternoon")
    }

}
