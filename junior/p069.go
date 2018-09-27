// 2019-09-12
// Reference: https://www.geeksforgeeks.org/eulers-totient-function-for-
// all-numbers-smaller-than-or-equal-to-n/
package main

import (
    "fmt"
    // "time"
)

func main() {
    // 1) Create an array phi[1..n] to store Φ values of all numbers
    //    from 1 to n.

    // 2) Initialize all values such that phi[i] stores i.  This
    //    initialization serves two purposes.
    //    a) To check if phi[i] is already evaluated or not. Note that
    //       the maximum possible phi value of a number i is i-1.
    //    b) To initialize phi[i] as i is a multiple in above product
    //       formula.

    // 3) Run a loop for p = 2 to n
    //     a) If phi[p] is p, means p is not evaluated yet and p is a
    //        prime number (slimier to Sieve), otherwise phi[p] must
    //        have been updated in step 3.b
    //     b) Traverse through all multiples of p and update all
    //        multiples of p by multiplying with (1-1/p).

    // 4) Run a loop from i = 1 to n and print all Ph[i] values.

    // start := time.Now()

    n := int(1e6)
    phi := make([]int, n+1)

    for i := 0; i <= n; i++ {
        phi[i] = i
    }

    for i := 2; i <= n; i++ {
        if phi[i] == i {
            for k := 1; k*i <= n; k++ {
                phi[k*i] = phi[k*i] / i * (i-1)
            }
        }
    }

    m := 1.0
    c := 1
    for i := 2; i <= n; i++ {
        mcurr := float64(i) / float64(phi[i])
        if mcurr > m {
            m = mcurr
            c = i
        }
    }

    // t := time.Now()

    fmt.Println(c)
    // fmt.Println(t.Sub(start))
}