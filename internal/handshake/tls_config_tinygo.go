//go:build tinygo

package handshake

import (
	"crypto/tls"
	"net"
)

func setupConfigForServer(conf *tls.Config, _, _ net.Addr) *tls.Config {
	return conf
}
