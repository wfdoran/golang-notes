package main
import "fmt"

func gen(a chan int) {
    for i := 0; i < 20; i++ {
        a <- i
    }
    close(a)
}

func mapp(in chan int, out chan int) {
    for x := range in {
        out <- x * x
    }
    close(out)
}

func filter(in chan int, out chan int) {
    for x := range in {
        if x % 2 == 0 {
            out <- x
        }
    }
    close(out)
}

func reduce(a chan int) int {
    rv := 0
    for x := range(a) {
        rv += x
    }
    return rv
}


func main() {
    a := make(chan int)
    b := make(chan int)
    c := make(chan int)
    
    go gen(a)
    go mapp(a, b)
    go filter(b, c)
    
    x := reduce(c)
    fmt.Println(x)
}