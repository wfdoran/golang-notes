package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
)

type R struct {
    C float32
}

func main() {
    fmt.Println("start client");
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Connection error", err)
    }
    encoder := gob.NewEncoder(conn)
    r := &R{3.5}
    encoder.Encode(r)
    conn.Close()
    fmt.Println("done");
}
