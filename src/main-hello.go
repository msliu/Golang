package main

// The import path follows $GOPATH to reference. For those custom packages, we
// should put into /lib for better reference.
import (
	"lib"
)

func main() {
  hello.PrintHello()
}