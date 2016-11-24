// 2016-11-22
package main

import (
    "fmt"
    // m"math"
)

func main() {

    // beginPrime := 673
    // for {
    //     beginPrime += 2
    //     if beginPrime % 3 != 0 && beginPrime % 7 != 0 &&
    //     isPrime(beginPrime) && isPrime(conc(673, beginPrime)) && isPrime(conc(beginPrime, 673)) &&
    //     isPrime(conc(109, beginPrime)) && isPrime(conc(beginPrime, 109)) &&
    //     isPrime(conc(7, beginPrime)) && isPrime(conc(beginPrime, 7)) &&
    //     isPrime(conc(3, beginPrime)) && isPrime(conc(beginPrime, 3)) {
    //         break
    //     }
    //     // fmt.Println(beginPrime)
    // }
    a := primeGenerate(30)
    for i := 0; i < len(a); i++ {
        if a[i] {
            fmt.Println(i)
        }
    }
}

// func conc(a, b int) int {
//     pow := 10
//     for b > pow {
//         pow *= 10
//     }
//     return a*pow + b
// }

// Sieve of Eratosthenes Algorithm
// https://primes.utm.edu/glossary/page.php?sort=SieveOfEratosthenes
func primeGenerate(under int) [under]bool {
    var a = [under]bool
    a[0] = false
    for i := 2; i < under; i++ {
        a[i] := true
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