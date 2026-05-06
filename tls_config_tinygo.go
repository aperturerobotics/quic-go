//go:build tinygo

package quic

import "crypto/tls"

func cloneTLSConfig(conf *tls.Config) *tls.Config {
	return conf
}
