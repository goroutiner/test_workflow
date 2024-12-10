package main

import (
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(10)

    for i := 0; i < 10; i++ {
        go func() {
            i++
            wg.Done()
        }()
    }

    wg.Wait()
}

func MaxInt(a, b int) int {
    if a >= b {
        return a
    }

    return b
}