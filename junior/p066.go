// 2016-12-11
// Reference: https://en.wikipedia.org/wiki/Pell%27s_equation
package main

import (
    "fmt"
    m"math"
)

func main() {
    D := 313.0
    sqrtD := m.Sqrt(D)
    y := 1.0
    x := 1.0
    found := false
    for ; !found; y++ {
        for x = m.Floor(y*sqrtD); x/y <= sqrtD; x++ {
            if Round(x*x - D*y*y) == 1 {
                found = true
                break
            }
        }
        if int(y) % 10000000000000 == 0 {
            fmt.Println("Million:", y)
        }
    }
    fmt.Println(found, int(x), int(y))
}

func Round(val float64) int {
    if val < 0 {
        return int(val-0.5)
    }
    return int(val+0.5)
}