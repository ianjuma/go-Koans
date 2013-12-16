package main

import (
    "fmt"
    "time"
)

var age int = 0

func worker(num []float64) (r float64) {
    panic("Not Implemented")
}

func realWorker(num []int) int {
    fmt.Println("Am busy waiting")
    fmt.Println(num)
    time.Sleep(5000)
    return num[1] / len(num)
}


func main() {

    //closure func
    sumAges := func(someArray []int) (total int) {

        for _, val := range(someArray) {
            total += val
        }

        return total
    }


    someData := []int {12, 23, 45, 657, 3252, 657, 45645, 543}

    fmt.Println("The sum of ages: ", sumAges(someData))


    total := int(0)
    //total := 0
    for _, val := range someData {
        total += val
    }

    fmt.Println("\nTotal is: ", total, "\n")
    // go worker(someData)
    go realWorker(someData)
    time.Sleep(1000) // wait for some data
    some_ret := th_()
    fmt.Println("Some ret ->", some_ret)

    multiple(12, 12.0)
    variadic(12, 34, 56)
    fmt.Println(variadic(3353, 543645, 6575, 4235, 4, 5, 6))
}

// specifying return type -- no type infrence
func th_() (int) {
    return 12
}

// no explicit return
func foo() (r int) {
    r = 12
    return
}

//returning multiple values
func multiple(x int, y float64) (int , float64) {
    return x, y
}


// Variadic func --takes inifinite args
func variadic(args ... int) int {
    total := 0
    for _, val := range args {
        total += val
    }
    return total
}
