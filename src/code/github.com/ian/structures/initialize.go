package main

import (
    "fmt"
    "math"
)

const myPie float64 = float64(3.14)
const s string = "constant"

const er = string("s")

func main() {

    // some math
    fmt.Println(myPie)
    fmt.Println(s)

    const n = 500000000
    const theD = 3e20 / n
    fmt.Println(int64(theD))

    fmt.Println(math.Sin(n))



    var x string = "Hey"
    fmt.Println(x)

    var a, b int = 1, 2
    fmt.Println(a, b)

    som := "String"
    fmt.Println(som)

    var d = true
    fmt.Println(d)

    var p = ":P"
    fmt.Println(p)

    var e int
    fmt.Println(e)

    week := int(7)
    // var week string = "Week"
    fmt.Println(week)

    months := string("Jan")
    fmt.Println(months)
}
