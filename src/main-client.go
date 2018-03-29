package main

import (
	"fmt"
	"net"
)

func sendData(conn net.Conn) {
  words := "Data Sent through TCP !!!!"
  conn.Write([]byte(words))
  fmt.Println("Data Sent..... ")
}

func main() {
  localHostAddr := "127.0.0.1:1024"
  tcpAddr, onError := net.ResolveTCPAddr("tcp4", localHostAddr)
  if onError != nil {
    fmt.Println("Fatal error: %s", onError.Error())
  }

  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  if err != nil {
    fmt.Println("Fatal error: %s", err.Error())
  }

  sendData(conn)
}