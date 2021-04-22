//+build !js

package quic

import (
	"net"
)

func wrapConn(pc net.PacketConn) (connection, error) {
	c, ok := pc.(OOBCapablePacketConn)
	if !ok {
		return &basicConn{PacketConn: pc}, nil
	}
	return newConn(c)
}
