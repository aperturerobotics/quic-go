//go:build !js
// +build !js

package quic

import (
	"net"
	"syscall"
)

func wrapConn(pc net.PacketConn) (connection, error) {
	conn, ok := pc.(interface {
		SyscallConn() (syscall.RawConn, error)
	})
	if ok {
		rawConn, err := conn.SyscallConn()
		if err != nil {
			return nil, err
		}
		err = setDF(rawConn)
		if err != nil {
			return nil, err
		}
	}
	c, ok := pc.(OOBCapablePacketConn)
	if !ok {
		return &basicConn{PacketConn: pc}, nil
	}
	return newConn(c)
}
