//go:build tinygo

package qerr

import "strconv"

func cryptoErrorMessage(e TransportErrorCode) string {
	return "tls alert " + strconv.FormatUint(uint64(e-0x100), 10)
}
