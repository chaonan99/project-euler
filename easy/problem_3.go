// 2016-11-20
package main

import (
    "fmt"
    m"math"
    "container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}


func main() {
    original := 600851475143
    for original % 2 == 0 {
        original /= 2
    }
    primePQ := &IntHeap{original}
    var p1, p2 int
    heap.Init(primePQ)
    for {
        p1, p2 = teaseApart(heap.Pop(primePQ).(int))
        if p1 == 1 {
            fmt.Println(p2)
            break
        }
        heap.Push(primePQ, p2)
        heap.Push(primePQ, p1)
    }
    heap.Push(primePQ, 3)
}

func teaseApart(ori int)(p1, p2 int) {
    div := int(m.Floor(m.Sqrt(float64(ori))))
    if div % 2 == 0 {
        div --
    }
    for ; ori % div != 0; div -= 2 {}
    p1 = div
    p2 = ori / div
    return
}
