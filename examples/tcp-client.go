package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
  for { 
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Send: ")
    text, _ := reader.ReadString('\n')
    conn, _ := net.Dial("tcp", "127.0.0.1:8080")
    fmt.Fprintf(conn, text + "\n")
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Receive: " + message)
    conn.Close()
  }
}