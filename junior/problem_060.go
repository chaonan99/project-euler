// 2016-11-22
package main

import (
    "fmt"
    // m"math"
)

func main() {
    // under := 3000000000 Not found under this
    a := primeGenerate(under)
    cp := 673
    for {
        cp += 2
        tmp1 := conc(673, cp)
        tmp2 := conc(cp, 673)
        if tmp1 > under || tmp2 > under {
            fmt.Println("Not found")
            break
        }
        if a[cp] && a[tmp1] && a[tmp2] &&
        a[conc(109, cp)] && a[conc(cp, 109)] &&
        a[conc(7, cp)] && a[conc(cp, 7)] &&
        a[conc(3, cp)] && a[conc(cp, 3)] {
            break
        }
    }
    fmt.Println(cp)
}

func conc(a, b int) int {
    pow := 10
    for b > pow {
        pow *= 10
    }
    return a*pow + b
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

// func isPrime(ori int) bool {
//     div := int(m.Floor(m.Sqrt(float64(ori))))
//     if div % 2 == 0 {
//         div --
//     }
//     for ; ori % div != 0; div -= 2 {}
//     return div == 1
// }