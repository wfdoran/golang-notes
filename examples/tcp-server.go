package main

import "net"
import "bufio"
import "strconv"
import "strings"

func main() {
  bins := 10
  vals := make([]int, bins)

  ln, _ := net.Listen("tcp", ":8080")
  for {
    conn, _ := ln.Accept()
    message1, _ := bufio.NewReader(conn).ReadString('\n')
    message2 := strings.TrimSpace(message1)
    
    i, err := strconv.Atoi(message2)
    
    if err != nil || i < 0 || i >= bins {
        conn.Write([]byte("Error\n"))
    } else {
        out := strconv.Itoa(vals[i])
        conn.Write([]byte(out + "\n"))
        vals[i]++
    }
    conn.Close()
  }
}