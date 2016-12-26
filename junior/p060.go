// 2016-12-07
package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    a := primeGenerate(83998399)
    l := listSet(a, 8400, 5)
    elapsed := time.Since(start)
    for _, r := range l {
        fmt.Println(r, "sum=", Sum(r))
    }
    fmt.Println("Solution 60 took ", elapsed)
}

func Sum(ori []int) (res int) {
    for _, o := range ori {
        res += o
    }
    return
}

func listSet(lut []bool, r int, aimLen int) [][]int {
    var resList [][]int
    for p := 3; p < r - aimLen*2; p += 2 {
        if lut[p] {
            res := generateSet(lut, lut[:r], []int{p}, aimLen)
            if len(res) > 0 {
                resList = append(resList, res)
            }
        }
    }
    return resList
}

func generateSet(lut []bool, rt []bool, seed []int, aimLen int) []int {
    if len(seed) >= aimLen {
        return seed
    }
    for p := seed[len(seed)-1] + 2; p < len(rt) ; p+=2 {
        if !rt[p] {
            continue
        }
        allPrime := true
        for _, si := range seed {
            if !lut[conc(si, p)] || !lut[conc(p, si)] {
                allPrime = false
                break
            }
        }
        if allPrime {
            res := generateSet(lut, rt, append(seed, p), aimLen)
            if len(res) > 0 {
                return res
            }
        }
    }
    return make([]int, 0)
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