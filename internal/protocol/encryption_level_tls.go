//go:build !tinygo

package protocol

import (
	"crypto/tls"
	"fmt"
)

func (e EncryptionLevel) ToTLSEncryptionLevel() tls.QUICEncryptionLevel {
	switch e {
	case EncryptionInitial:
		return tls.QUICEncryptionLevelInitial
	case EncryptionHandshake:
		return tls.QUICEncryptionLevelHandshake
	case Encryption1RTT:
		return tls.QUICEncryptionLevelApplication
	case Encryption0RTT:
		return tls.QUICEncryptionLevelEarly
	default:
		panic(fmt.Sprintf("unexpected encryption level: %s", e))
	}
}

func FromTLSEncryptionLevel(e tls.QUICEncryptionLevel) EncryptionLevel {
	switch e {
	case tls.QUICEncryptionLevelInitial:
		return EncryptionInitial
	case tls.QUICEncryptionLevelHandshake:
		return EncryptionHandshake
	case tls.QUICEncryptionLevelApplication:
		return Encryption1RTT
	case tls.QUICEncryptionLevelEarly:
		return Encryption0RTT
	default:
		panic(fmt.Sprintf("unexpect encryption level: %s", e))
	}
}
