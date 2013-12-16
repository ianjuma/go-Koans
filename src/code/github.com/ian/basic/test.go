package basic

import "fmt"


func get() {

    for i := 1; i < 100; i++ {
        if i % 3 == 0 {
            fmt.Println(i)
        }
    }
}
