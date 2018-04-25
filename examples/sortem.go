package main

import "fmt"
import "sort"

type Pair struct {
    X int
    Y float32
}

func makeLess(a []Pair) func(i, j int) bool {
    return func(i, j int) bool {
        if a[i].X != a[j].X {
            return a[i].X < a[j].X
        }
        return a[i].Y < a[j].Y
    }
}

func main() {
    n := 10
    p := make([]Pair, n)
    p[0].X = 1
    p[0].Y = 0.0
    
    for i := 1; i < n; i++ {
        p[i].X = (p[i-1].X + 73) % 100
        p[i].Y = p[i-1].Y + .636363636
        if p[i].Y > 1.0 {
            p[i].Y -= 1.0
        }
    }
    
    p_less := makeLess(p)
    sort.Slice(p, p_less)
        
    fmt.Println(p)
}