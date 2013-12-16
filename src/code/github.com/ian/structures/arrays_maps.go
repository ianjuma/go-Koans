package main

import "fmt"

func main() {

    //var todos []int
    todos := make([]int, 5) //slicing
    todos[1] = 12

    slice1 := append(todos, 1, 2, 3)

    fmt.Println(len(slice1)) //8

    fmt.Println(len(todos))
    fmt.Println("!")

    var ids[5] int

    x := [5]float64 {
        98,
        93,
        90,
        89,
        97,
    }

    fmt.Println(len(x))

    i := 0
    for i < len(ids) { // for i < 5 {}
        ids[i] = (i*90)
        i += 1
    }

    var total int = 0

    for _, value := range ids {
        total += value
    }

    fmt.Println(ids)
    fmt.Print("Size of array is -> ")
    fmt.Println(len(ids))
    fmt.Println(total)

}
