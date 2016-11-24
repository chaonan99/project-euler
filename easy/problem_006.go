// 2016-11-22
package main

import (
    "fmt"
)

func main() {
    sumDiff := 0
    for i := 0; i < 100; i++ {
        for j := i + 1; j <= 100; j++ {
            sumDiff += i*j
        }
    }
    fmt.Println(sumDiff*2)
}