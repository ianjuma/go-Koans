package main


import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type square struct {
    height, width float64
}

type circle struct {
    radius float64
}


func (s square) area() float64 {
    return s.height * s.width
}

func (s square) perim() float64 {
    return 2*s.height + 2*s.width
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius+c.radius
}

func measure(g geometry) {
    fmt.Println(g.area())
    fmt.Println(g.perim())
}


func main() {
    c := circle{radius: 7.0}
    s := square{height: 4.0, width: 4.0}

    measure(s)
    measure(c)
}
