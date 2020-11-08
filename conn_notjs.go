//+build !js

package quic

import (
	"net"

	"github.com/lucas-clemente/quic-go/internal/utils"
)

func wrapConn(pc net.PacketConn) (connection, error) {
	c, ok := pc.(OOBCapablePacketConn)
	if !ok {
		utils.NewDefaultLogger(nil).Infof("PacketConn is not a net.UDPConn. Disabling optimizations possible on UDP connections.")
		return &basicConn{PacketConn: pc}, nil
	}
	return newConn(c)
}
