package main

import "fmt"

type Person struct {
    name string
    age int
}

func main() {
    fmt.Println(Person{"Ian", 20})
    fmt.Println(Person{name: "Juma", age: 20})
    fmt.Println(Person{age: 21})

    fmt.Println(&Person{name: "Mark", age: 56})

    data := Person{name: "Ham", age: 12}
    fmt.Println(data.name)
    data.name = "Andrea"
    fmt.Println(data.name)

    data_ref := &data //pass by ref
    data_ref.name = "Mark"
    fmt.Println(data.name)

}
