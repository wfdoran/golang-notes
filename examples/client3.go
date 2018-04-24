package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
)

type P struct {
    A int
    B int
    C float32
}

func main() {
    fmt.Println("start client");
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Connection error", err)
    }
    encoder := gob.NewEncoder(conn)
    p := &P{1,2,3.5}
    encoder.Encode(p)
    conn.Close()
    fmt.Println("done");
}
