// 2016-12-04
package main

import (
    "fmt"
    // m"math"
)

func main() {
    tri := 1
    for i := 2; ; i++ {
        tri += i
        if nFactor(tri) > 500 {
            fmt.Println(tri)
            break
        }
    }
}

// Reference: http://www.cut-the-knot.org/blue/NumberOfFactors.shtml
func nFactor(s int) d int {
    
}