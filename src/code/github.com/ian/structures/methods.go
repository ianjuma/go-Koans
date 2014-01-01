// methods and member objects

package main

import (
    "fmt"
)

type rect struct {
    width, height int
}


func(r *rect) area() int {
    return r.width * r.height
}


func(r *rect) perim() int {
    return 2*r.width + 2*r.height
}


func main() {
    r := rect{width: 12, height: 2}

    fmt.Println("area", r.area())
    fmt.Println("Perim", r.perim())

    sr := &r
    fmt.Println("area2 ->", sr.area())
}
