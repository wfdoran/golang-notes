package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
)

type Q struct {
    A int
    B int
}

func main() {
    fmt.Println("start client");
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Connection error", err)
    }
    encoder := gob.NewEncoder(conn)
    q := &Q{1, 2}
    encoder.Encode(q)
    conn.Close()
    fmt.Println("done");
}
