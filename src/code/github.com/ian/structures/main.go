package main

import "fmt"

func hello(some string) string {
    // defn definition with a retrun value
    return some+" Its Me!"
}

func foo() {
    fmt.Println("Foo :: foo()")
}


func main() {
    foo()
    fmt.Println(hello("Juma"))

    var myName string = "Juma"
    otherName := "Wafula"

    fmt.Println(otherName + " \t "+ " \t<----")

    fmt.Println(myName)

    var worker_n = 12
    workers := 1

    fmt.Println(1)
    fmt.Println(workers)
    fmt.Println(worker_n)

    fmt.Print("\n")
    fmt.Println("Even Odd Number Test")

    i := 1
    for i <= 10 {
        fmt.Print(i)
        fmt.Print(" ")

        if i % 2 == 0 {
            fmt.Println("Even")
        } else {
            fmt.Println("Odd")
        }

        i += 1
    }
}
