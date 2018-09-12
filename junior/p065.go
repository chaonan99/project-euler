// 2019-09-12
package main

import (
    "fmt"
    "math/big"
)

func main() {
    // Generate sequence
    seq := make([]int, 98)
    for i := 0; i < 98; i++ {
        if i % 3 == 1 {
            seq[i] = (i + 2) / 3 * 2
        } else {
            seq[i] = 1
        }
    }

    p := new(big.Rat).SetInt64(1)
    c := new(big.Rat)
    for i := 97; i >= 0; i-- {
        c.SetInt64(int64(seq[i]))
        p.Add(p, c).Inv(p)
    }
    p.Add(p, new(big.Rat).SetInt64(int64(2)))
    numstr := p.Num().String()

    sum := 0
    for _, ch := range numstr {
        sum += (int(ch) - 48)
    }

    fmt.Println(sum)
}
