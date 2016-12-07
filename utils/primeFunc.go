package main
import "fmt"
import "time"
import m"math"


func main() {
    start := time.Now()

    // for i := 0; i < 10; i++ {
    //     _ = SieveEra(10000000)
    // }
    // _ = SieveEraFast(1000000000)
    fmt.Println(IsPrimeTD(26033))
    fmt.Println(IsPrimeTD(260337))
    fmt.Println(IsPrimeTD(260333))
    fmt.Println(IsPrimeTD(26033109))
    fmt.Println(IsPrimeTD(26033673))
    fmt.Println(IsPrimeTD(726033))
    fmt.Println(IsPrimeTD(326033))
    fmt.Println(IsPrimeTD(10926033))
    fmt.Println(IsPrimeTD(67326033))
    // fmt.Printf("%v\n", FactorizeTD(10212305183))
    elapsed := time.Since(start)
    fmt.Printf("Sieve of Eratosthenes took %s", elapsed)
}


// Sieve of Eratosthenes Algorithm
// https://primes.utm.edu/glossary/page.php?sort=SieveOfEratosthenes
func SieveEra(under int) []bool {
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
func SieveEraFast(under int) []bool {
    m := (under - 1) / 2
    a := make([]bool, m)
    for i := 0; i < m; i++ {
        a[i] = true
    }
    i := 0
    p := 3
    // fmt.Println(2)
    count := 1
    for p*p <= under {
        if a[i] {
            // fmt.Println(p)
            count++

            j := (p*p - 3) / 2
            for j < m {
                a[j] = false
                j = j+p
            }
        }
        i++
        p += 2
    }
    for i < m {
        if a[i] {
            // fmt.Println(p)
            count++
        }
        i++
        p += 2
    }
    fmt.Println("Prime found: ", count)
    return a
}


// Trial Division for checking prime
func IsPrimeTD(ori int) bool {
    if ori % 2 == 0 {
        return false
    }
    limit := int(m.Floor(m.Sqrt(float64(ori))))
    div := 3
    for ; ori % div != 0 && div <= limit; div += 2 {}
    return div > limit
}


// Factorization using trial division.
func FactorizeTD(ori int) (fcs []int) {
    // var fcs []int
    for ; ori % 2 == 0; ori /= 2 {
        fcs = append(fcs, 2)
    }
    for fc := 3; ori > 1 && ori > fc*fc; fc += 2 {
        for ;ori % fc == 0; ori /= fc {
            fcs = append(fcs, fc)
        }
    }
    if ori > 1 {
        fcs = append(fcs, ori)
    }
    return
}


func GCD(a, b int) int {
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