package main

import (
    "fmt"
    "errors"
)

type argError struct {
    arg int
    argProb string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.argProb)
}

func foo(arg int) (int, error) {
    if arg == 13 {
        return -1, errors.New("I don't like 13")
    }

    return arg, nil
}

func main() {
    for _, i := range []int{12, 13} {
        if r, e := foo(i); e != nil {
            fmt.Println("Foo failed", i, r)
        } else {
            fmt.Println("F1 worked()", r)
        }
    }

    foo(13)
}
