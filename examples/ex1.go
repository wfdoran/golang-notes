package main

import "fmt"
import "./int128"

func main() {
    a := int128.Init(10000000)
    b := int128.Init(20000)
    c := int128.Add(a,b)
    fmt.Println(c)
}
