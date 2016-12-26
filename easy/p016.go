// 2016-12-26

package main

import (
    "fmt"
)

func main() {
    sumSlice := []byte{2}
    for i := 1; i < 1000; i ++ {
        sumSlice = multi2(sumSlice)
    }
    fmt.Println(sumEle(sumSlice))
}

func multi2(ori []byte) []byte {
    up := byte(0)
    for i := 0; i < len(ori); i ++ {
        tmp := ori[i]*2 + up
        ori[i] = tmp % 10
        up = tmp / 10
    }
    if up > 0 {
        ori = append(ori, up)
    }
    return ori
}

func sumEle(ori []byte) (sum int) {
    for i := 0; i < len(ori); i ++ {
        sum += int(ori[i])
    }
    return
}