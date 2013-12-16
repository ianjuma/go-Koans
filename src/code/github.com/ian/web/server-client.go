package main

import (
    "encoding/json"
    "fmt"
    "net"
)

func server() {
    ln, err := net.Listen("", ":8000")
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        c, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        go handleServerConnection(c)
    }
}

func handleServerConnection(c net.Conn) {
    var msg string
    err := json.NewDecoder(c).Decode(&msg)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Received", msg)
    }

    c.Close()
}


func client() {
    c, err := net.Dial("tcp", "127.0.0.1:8000")
    if err != nil {
        fmt.Println(err)
        return
    }

    msg := "Hello World"
    fmt.Println("Sending", msg)

    err = json.NewEncoder(c).Encode(msg)
    if err != nil {
        fmt.Println(err)
    }

    c.Close()
}

func main() {
    go server()
    go client()

    var input string
    fmt.Scanln(&input)
}
