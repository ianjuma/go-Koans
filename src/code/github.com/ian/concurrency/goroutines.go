// +build main1
package main

import (
    "fmt"
    "time"
)


// events + run
func simulateEvent(name string, timeInSecs float64) {
    fmt.Println("Started", name, ": Should take", timeInSecs, "seconds.")
    time.Sleep(timeInSecs * 2e3)
    fmt.Println("Finished", name)
}


// goroutines launch + main on separate goroutine
func main() {
    go simulateEvent("100 Sprint", 30)
    go simulateEvent("Long jump", 12)
    go simulateEvent("High Jump", 2)

    time.Sleep(17 * 1e9) // wait for go-routines to finish
    // increase time-stamp to take longer, to see the events
}
