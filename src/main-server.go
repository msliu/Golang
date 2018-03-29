package main

import (
  "net"
  "log"
  "fmt"
)

// Refers to https://golang.org/pkg/net/#Listener
func main() {
	netListener, onError := net.Listen("tcp", "localhost:1024")
	if onError != nil {
      log.Fatal(onError)
	}
	// Refers to https://golang.org/doc/effective_go.html#defer
	defer netListener.Close()
    for {
      // Wait for a connection.
      fmt.Println("Waiting for clients ... ")
      conn, onError := netListener.Accept()
      if onError != nil {
      	continue
      }
      handleConnection(conn)
    }  
}

func handleConnection(c net.Conn) {
  buffer := make([]byte, 128)
  for {
    n, onError := c.Read(buffer)
    if onError != nil {
      return
    }
    log.Println(c.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
  }
}