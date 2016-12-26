// 2016-11-22
package main

import (
    "fmt"
    m"math"
)

func main() {
    primeFind := 3
    counter := 2
    for counter < 10001 {
        primeFind += 2
        if isPrime(primeFind) {
            counter += 1
        }
    }
    fmt.Println(primeFind)
}

func isPrime(ori int) bool {
    div := int(m.Floor(m.Sqrt(float64(ori))))
    if div % 2 == 0 {
        div --
    }
    for ; ori % div != 0; div -= 2 {}
    return div == 1
}