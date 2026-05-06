//go:build !tinygo

package quic

import "net"

func listenUDPAddr(addr *net.UDPAddr) (*net.UDPConn, error) {
	return net.ListenUDP("udp", addr)
}

var _ OOBCapablePacketConn = &net.UDPConn{}
