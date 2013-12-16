package main

import "fmt"

func main() {
    sl := make([]int, 5)

    fmt.Println(sl, "Of len -> ", len(sl))
    sl[0] = 12
    sl[1] = 2
    sl[2] = 12
    sl[3] = 23
    sl[4] = 43

    fmt.Println(sl)
    fmt.Println("Slice 1", sl[1])
    sl = append(sl, 12, 34)
    fmt.Println("-->", sl)
    fmt.Println("Len ->", len(sl))

    sl2 := make([]int, 5)
    sl3 := []int{12, 34, 45, 56, 67}
    fmt.Println("Funny slice", sl3[:4])


    copy(sl2, sl)
    fmt.Println(sl2)

    someSlice := sl2[1:3]
    someSlice2 := sl2[3:] // 3 onwards
    someSlice3 := sl2[:4]
    fmt.Println(someSlice)
    fmt.Println(someSlice2)
    fmt.Println(someSlice3)
}
