// 2016-11-22
package main

import (
    "fmt"
    // m"math"
)

// func main() {
//     fmt.Println("begin")
//     find := false
//     for i := 3; i <= 333 && !find; i++ {
//         for j := i+1; j <= int(float64(1000-i)/2.0) && !find; j++ {
//             if i*i + j*j - (1000-i-j)*(1000-i-j) == 0 {
//                 fmt.Println(i,j)
//                 find = true
//             }
//         }
//     }
//     fmt.Println("end")
// }

// Parameterize Pythagorean triplet. See 009_overview.
func main() {
    s := 1000
    s2 := s/2
    for m := 2; m*m < s2; m++ {
        if s2 % m == 0 {
            sm := s2/m
            var k int
            if m % 2 == 1 {
                k = m+2
            } else {
                k = m+1
            }
            for ; k <= sm && k < 2*m; k += 2 {
                if sm % k == 0 && gcd(m, k) == 1 {
                    d := s2 / k / m
                    n := k - m
                    a := d*(m*m-n*n)
                    b := d*2*m*n
                    c := d*(m*m+n*n)
                    fmt.Println(a,b,c)
                }
            }
        }
    }
}

// Function to find GCD
func gcd(a, b int) int {
    if a == b {
        return b
    }
    if a > b {
        a = a % b
        if a == 0 {
            return b
        }
    }
    for {
        b = b % a
        if b == 0 {
            return a
        }
        a = a % b
        if a == 0 {
            return b
        }
    }
}