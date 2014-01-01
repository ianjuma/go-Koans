// simple key value pairs

package main

import "fmt"

func main() {

    x := make(map[string]int)

    x["Ian"] = 20
    x["Jay"] = 12
    x["Alex"] = 17

    fmt.Println(x["Ian"], x["Jay"], x["Alex"])

    name, age := x["Ian"]
    fmt.Println(name, age)

    delete(x, "Ian")

    y := map[string]string{"FName": "Ian", "LName": "Juma"}
    fmt.Println(y["FName"])

    _, pr := y["YName"]
    fmt.Println("pr ->", pr)


    fmt.Println(x["Ian"], x["Jay"], x["Alex"])
}
