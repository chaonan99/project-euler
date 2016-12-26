// 2016-12-06
package main

import (
    "fmt"
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
func nFactor(ori int) int {
    var pns []int
    flag := true
    for ; ori % 2 == 0; ori /= 2 {
        if flag {
            flag = false
            pns = append(pns, 1)
        }
        pns[len(pns)-1]++
    }
    flag = true
    for fc := 3; ori > 1 && ori > fc*fc; fc += 2 {
        for ;ori % fc == 0; ori /= fc {
            if flag {
                flag = false
                pns = append(pns, 1)
            }
            pns[len(pns)-1]++
        }
        flag = true
    }
    if ori > 1 {
        pns = append(pns, 2)
    }
    nfactor := 1
    for _, pn := range pns {
        nfactor *= pn
    }
    return nfactor
}