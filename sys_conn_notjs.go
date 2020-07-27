//go:build !js
// +build !js

package quic

import (
	"net"
	"syscall"
)

func wrapConn(pc net.PacketConn) (rawConn, error) {
	conn, ok := pc.(interface {
		SyscallConn() (syscall.RawConn, error)
	})
	var supportsDF bool
	if ok {
		rawConn, err := conn.SyscallConn()
		if err != nil {
			return nil, err
		}

		if _, ok := pc.LocalAddr().(*net.UDPAddr); ok {
			// Only set DF on sockets that we expect to be able to handle that configuration.
			var err error
			supportsDF, err = setDF(rawConn)
			if err != nil {
				return nil, err
			}
		}
	}
	c, ok := pc.(OOBCapablePacketConn)
	if !ok {
		return &basicConn{PacketConn: pc, supportsDF: supportsDF}, nil
	}
	return newConn(c, supportsDF)
}
