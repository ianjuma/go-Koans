package main

import "fmt"

func main() {
    a := []int{12, 34, 5} // ! var a []int{12,34, 5} -> wont work
    fmt.Println(a)

    var b [5]int
    c := []int{12, 23} // c := [5]int{12, 23, 45, 56, 67}

    b[0] = 12
    c[0] = 20
    fmt.Println(b)
    fmt.Println(c, "Len of c ->", len(c))

    // 2-d arrays -> nd arrays
    var nd [3][2]int
    for val := 0; val < 3; val ++ {
        for innerVal := 0; innerVal < 2; innerVal++ {
            nd[val][innerVal] = innerVal + val
        }
    }
    fmt.Println(nd)

}
