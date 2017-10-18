// A simple proxy that multiplexes a unix socket connection
//
// Copyright 2017 HyperHQ
package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/hashicorp/yamux"
)

func server(channel string) error {
	// just remove old ones for testing
	os.Remove(channel)
	l, err := net.Listen("unix", channel)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		session, err := yamux.Server(conn, nil)
		if err != nil {
			return err
		}

		for {
			stream, err := session.Accept()
			if err != nil {
				fmt.Println("stream accept failed: ", err)
				break
			}
			fmt.Println("New yamux connection")
			go io.Copy(os.Stdout, stream)
		}
	}

}
func main() {
	vmChannel := "/tmp/target.sock"

	server(vmChannel)
}
