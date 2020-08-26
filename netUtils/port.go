package netUtils

import (
	"fmt"
	"net"
)

func PickFreePort() (string, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	// l, err := net.Listen("tcp", "")
	if err != nil {
		return "", err
	}
	defer l.Close()
	addr := l.Addr().String()
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}
	return port, nil
}
