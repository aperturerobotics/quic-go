//go:build !tinygo

package qerr

import "crypto/tls"

func cryptoErrorMessage(e TransportErrorCode) string {
	return tls.AlertError(e - 0x100).Error()
}
