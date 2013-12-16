package main

import "fmt"

func main() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i += 1
    }

    for x:=0; x < 10; x++ {
        fmt.Println(x)
        x += 1
    }

    for count := 0; count < 10; count++ {
        fmt.Println(count)
    }

    for {
        fmt.Println("Infinite-Loop")
        break
    }


}
