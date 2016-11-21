// 2016-11-19
package main

import (
    "fmt"
)

// Solution 1
// func main() {
//     upBound := 1000
//     sumResult := 0
//     for x := 0; x < upBound; x++ {
//         if x % 3 == 0 || x % 5 == 0 {
//             sumResult += x
//         }
//     }
//     fmt.Println(sumResult)
// }


// Solution 2
func main() {
    upBound := 1000
    result := sumDividedBy(3, upBound) + sumDividedBy(5, upBound) - sumDividedBy(15, upBound)
    fmt.Println(result)
}

func sumDividedBy (div, upBound int) (sum int) {
    numberOfAdds := (upBound - 1) / div
    sum = (1 + numberOfAdds) * numberOfAdds / 2 * div
    return
}