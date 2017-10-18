// A simple proxy that multiplexes a unix socket connection
//
// Copyright 2017 HyperHQ
package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/golang/glog"
	"github.com/hashicorp/yamux"
)

// @channel is the unix socket address we want to multiplex
func serv(channel, proto, addr string) error {
	// yamux connection
	servConn, err := net.Dial("unix", channel)
	if err != nil {
		glog.Errorf("fail to dial channel(%s): %s", channel, err.Error())
		return err
	}
	session, err := yamux.Client(servConn, nil)
	if err != nil {
		glog.Errorf("fail to create yamux client: %s", err.Error())
		return err
	}

	// serving connection
	l, err := net.Listen(proto, addr)
	if err != nil {
		glog.Errorf("fail to listen on %s:%s: %s", proto, addr, err.Error())
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			glog.Errorf("fail to accept new connection: %s", err.Error())
			return err
		}
		stream, err := session.Open()
		if err != nil {
			glog.Errorf("fail to open yamux stream: %s", err.Error())
			return err
		}
		go io.Copy(conn, stream)
		go io.Copy(stream, conn)
	}
}

func main() {
	pid := os.Getpid()
	vmChannel := "/tmp/testvm.sock"
	proxyAddr := fmt.Sprintf("/tmp/proxy-%d.sock", pid)

	serv(vmChannel, "unix", proxyAddr)
}