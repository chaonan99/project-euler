// 2016-11-19
package main

import (
    "fmt"
)

// func main() {
//     n1 := 1
//     n2 := 2
//     sumResult := 0
//     for n2 <= 4000000 {
//         if n2 % 2 == 0 {
//             sumResult += n2
//         }
//         n3 := n1 + n2
//         n1 = n2
//         n2 = n3
//     }
//     fmt.Println(sumResult)
// }

// Better solution
func main() {
    n1 := 2
    n2 := 8
    sumResult := 2
    for n2 <= 4000000 {
        sumResult += n2
        n3 := n1 + 4 * n2
        n1 = n2
        n2 = n3
    }
    fmt.Println(sumResult)
}
