// 2016-12-04
package main

import (
    "fmt"
    // m"math"
)

func main() {
    prime_list := primeGenerate(2000000)
    sum := 0
    for i := 1; i < 2000000; i++ {
        if prime_list[i] {
            sum += i
        }
    }
    fmt.Println(sum)
}

// Sieve of Eratosthenes Algorithm
// https://primes.utm.edu/glossary/page.php?sort=SieveOfEratosthenes
func primeGenerate(under int) []bool {
    a := make([]bool, under)
    for i := 2; i < under; i++ {
        a[i] = true
    }
    for p := 2; p*p <= under; p++ {
        if a[p] {
            j := p*p
            for j < under {
                a[j] = false
                j = j+p
            }
        }
    }
    return a
}