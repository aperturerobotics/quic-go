//+build js

package quic

import "net"

func wrapConn(pc net.PacketConn) (connection, error) {
	// no ECN available in JS
	return &basicConn{PacketConn: pc}, nil
}
