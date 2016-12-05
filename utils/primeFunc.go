package main
import "fmt"
import "time"


func main() {
    start := time.Now()

    // for i := 0; i < 10; i++ {
    //     _ = sieveEra(10000000)
    // }
    _ = sieveEraFast(10000000)
    elapsed := time.Since(start)
    fmt.Printf("Sieve of Eratosthenes (10 times) took %s", elapsed)
}


// Sieve of Eratosthenes Algorithm
// https://primes.utm.edu/glossary/page.php?sort=SieveOfEratosthenes
func sieveEra(under int) []bool {
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


// Optimized
// https://programmingpraxis.files.wordpress.com/2012/09/primenumbers.pdf
// Exclude even numbers to save space.
func sieveEraFast(under int) []bool {
    m := (under - 1) / 2
    a := make([]bool, m)
    for i := 0; i < m; i++ {
        a[i] = true
    }
    i := 0
    p := 3
    // fmt.Println(2)
    for p*p <= under {
        if a[i] {
            // fmt.Println(p)

            j := (p*p - 3) / 2
            for j < m {
                a[j] = false
                j = j+p
            }
        }
        i++
        p += 2
    }
    // for i < m {
    //     if a[i] {
    //         fmt.Println(p)
    //     }
    //     i++
    //     p += 2
    // }
    return a
}


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