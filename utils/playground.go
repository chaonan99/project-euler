package main

import (
    "fmt"
)

func main() {
    ori := []byte{2}
    ori[1] = 3
    fmt.Println(ori[1])
}