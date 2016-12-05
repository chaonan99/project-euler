package main

import (
    "fmt"
)

func main() {
    i := 0
    p := 2
    for  p < 50 {
        fmt.Println(i)
        i++
        p += 2
    }
}