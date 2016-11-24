// 2016-11-21
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // foundFlag := false
    currentMax := 0
    for i := 999; i >= 100; i-- {
        for j := 999; j >= i; j-- {
            tempProduct := i*j
            if isPalindromic(tempProduct) {
                if tempProduct > currentMax {
                    currentMax = tempProduct
                }
                // fmt.Println(i*j, i, j)
                // foundFlag = true
                // break
            }
        }
        // if foundFlag {
        //     break
        // }
    }
    fmt.Println(currentMax)
}

func isPalindromic (ori int) bool {
    oris := strconv.Itoa(ori)
    n := len(oris)
    for i := 0; i < n / 2; i++ {
        if oris[i] != oris[n-1-i] {
            return false
        }
    }
    return true
}