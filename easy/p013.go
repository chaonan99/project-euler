// 2016-12-26

package main

import (
    "io/ioutil"
    "fmt"
    "strings"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// int32 range: -2147483648 through 2147483647, not sufficient for the program use
// int64 range:  -9223372036854775808 through 9223372036854775807, sufficient
func main() {
    dat, err := ioutil.ReadFile("problem_013_input.txt")
    check(err)
    numbers := strings.Split(string(dat), "\r\n")
    // fmt.Println(len(numbers), len(numbers[0]), numbers[0][40:50])
    currSum := int64(0)
    for i := len(numbers[0]); i > 0; i -= 10 {
        tmpSS := strconv.FormatInt(currSum, 10)
        if len(tmpSS) > 10 {
            tmp, err := strconv.ParseInt(tmpSS[0:len(tmpSS)-10], 10, 64)
            check(err)
            currSum = tmp
        } else {
            currSum = 0
        }
        for j := 0; j < len(numbers); j++ {
            tmp, err := strconv.ParseInt(numbers[j][i-10:i], 10, 64)
            check(err)
            currSum += tmp
        }
    }
    fmt.Println(strconv.FormatInt(currSum, 10)[0:10])
}
