// 2019-09-11
package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    m"math"
    "strconv"
    "github.com/golang/geo/r2"
    "github.com/golang/geo/s1"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    dat, err := ioutil.ReadFile("../input/p102_triangles.txt")
    check(err)
    datstr := string(dat)
    lines := strings.Split(datstr, "\n")

    count := 0
    for _, line := range lines {
        points := strings.Split(line, ",")
        if len(points) == 6 {
            allpoints := make([]r2.Point, 3)
            for j := 0; j < 3; j++ {
                pnumx, err := strconv.Atoi(points[j*2])
                check(err)
                pnumy, err := strconv.Atoi(points[j*2 + 1])
                check(err)

                allpoints[j] = r2.Point{float64(pnumx), float64(pnumy)}
            }

            allnorms := make([]float64, 3)
            for j := 0; j < 3; j++ {
                allnorms[j] = allpoints[j].Norm()
            }

            sumangle := 0.0
            for j := 0; j < 3; j++ {
                i := (j + 1) % 3
                cos := allpoints[i].Dot(allpoints[j]) / allnorms[i] / allnorms[j]
                sumangle += m.Acos(cos)
            }
            if m.Abs(s1.Angle(sumangle).Degrees() - 360.0) < 1e-5 {
                count++
            }
        }
    }
    fmt.Println(count)
}

