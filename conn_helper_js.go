// +build js

package quic

const msgTypeIPTOS = 0x1

const (
	ipv4RECVPKTINFO = 0x8
	ipv6RECVPKTINFO = 0x31
)

const (
	msgTypeIPv4PKTINFO = 0x8
	msgTypeIPv6PKTINFO = 0x32
)

const batchSize = 8
