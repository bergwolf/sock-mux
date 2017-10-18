// A simple proxy that multiplexes a unix socket connection
//
// Copyright 2017 HyperHQ
package main

import (
	"fmt"
	"net"
)

func main() {
	buf := []byte("hello proxy\n")
	proxyAddr := "/tmp/proxy.sock"

	conn, err := net.Dial("unix", proxyAddr)
	if err != nil {
		fmt.Println("dial failed: ", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println("write failed: ", err)
	}
}
