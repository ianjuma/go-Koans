package main

import "fmt"

func main() {
    nums := []int{1, 23, 45, 76, 3, 90, 5}
    sum := 0

    for _, num := range nums {
        sum += num
    }

    fmt.Println(sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i, "Position:", i+1)
        }
    }

    kvs := map[string]string{"a": "Apple", "b": "Banana"}
    for k, v := range kvs {
        fmt.Println("Key -> Value", k, v)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }

}
