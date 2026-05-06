//go:build tinygo

package quic

import (
	"errors"
	"net"
)

var errTinyGoUDPUnsupported = errors.New("quic-go: UDP sockets are unsupported by TinyGo")

func listenUDPAddr(*net.UDPAddr) (*net.UDPConn, error) {
	return nil, errTinyGoUDPUnsupported
}
