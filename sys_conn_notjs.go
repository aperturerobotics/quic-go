//go:build !js
// +build !js

package quic

import (
	"net"
	"syscall"
)

func wrapConn(pc net.PacketConn) (rawConn, error) {
	if err := setReceiveBuffer(pc); err != nil {
		/*
			if !strings.Contains(err.Error(), "use of closed network connection") {
				setBufferWarningOnce.Do(func() {
					if disable, _ := strconv.ParseBool(os.Getenv("QUIC_GO_DISABLE_RECEIVE_BUFFER_WARNING")); disable {
						return
					}
					log.Printf("%s. See https://github.com/quic-go/quic-go/wiki/UDP-Buffer-Sizes for details.", err)
				})
			}
		*/
	}
	if err := setSendBuffer(pc); err != nil {
		/*
			if !strings.Contains(err.Error(), "use of closed network connection") {
				setBufferWarningOnce.Do(func() {
					if disable, _ := strconv.ParseBool(os.Getenv("QUIC_GO_DISABLE_RECEIVE_BUFFER_WARNING")); disable {
						return
					}
					log.Printf("%s. See https://github.com/quic-go/quic-go/wiki/UDP-Buffer-Sizes for details.", err)
				})
			}
		*/
	}

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
		// utils.DefaultLogger.Infof("PacketConn is not a net.UDPConn. Disabling optimizations possible on UDP connections.")
		return &basicConn{PacketConn: pc, supportsDF: supportsDF}, nil
	}
	return newConn(c, supportsDF)
}
