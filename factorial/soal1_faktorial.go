package main

import (
    "fmt"
    "math"
)

func faktorial(n int) float64 {
    if n < 0 {
        return 0
    }
    if n == 0 {
        return 1
    }
    hasil := 1.0
    for i := 1; i <= n; i++ {
        hasil *= float64(i)
    }
    return hasil
}

func f(n int) int {
    fakt := faktorial(n)
    pangkatDua := math.Pow(2, float64(n))
    hasil := fakt / pangkatDua
    return int(math.Ceil(hasil))
}

func main() {
    for _, n := range []int{0, 4, 5, 7} {
        fmt.Printf("f(%d) = %d\n", n, f(n))
    }
}