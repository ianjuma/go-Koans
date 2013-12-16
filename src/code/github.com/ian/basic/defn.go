package basic

import "fmt"

func soMe() {
    num := make()
    fmt.Println(num())
    fmt.Println(num())

    for val := 0; val < 10; val++ {
        fmt.Println(num())
    }
}


func make() func() uint {
    i := uint(2)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}
