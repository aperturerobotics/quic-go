//go:build tinygo

package handshake

import (
	"context"
	"crypto/tls"
	"errors"
	"net"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
	"github.com/quic-go/quic-go/internal/wire"
	"github.com/quic-go/quic-go/qlogwriter"
)

var errTinyGoQUICUnsupported = errors.New("quic-go: QUIC TLS is unsupported by TinyGo")

type quicVersionContextKey struct{}

var QUICVersionContextKey = &quicVersionContextKey{}

type cryptoSetup struct {
	initialOpener LongHeaderOpener
	initialSealer LongHeaderSealer
}

var _ CryptoSetup = &cryptoSetup{}

func NewCryptoSetupClient(
	connID protocol.ConnectionID,
	_ *wire.TransportParameters,
	_ *tls.Config,
	_ bool,
	_ *utils.RTTStats,
	_ qlogwriter.Recorder,
	_ utils.Logger,
	version protocol.Version,
) CryptoSetup {
	initialSealer, initialOpener := NewInitialAEAD(connID, protocol.PerspectiveClient, version)
	return &cryptoSetup{initialOpener: initialOpener, initialSealer: initialSealer}
}

func NewCryptoSetupServer(
	connID protocol.ConnectionID,
	_, _ net.Addr,
	_ *wire.TransportParameters,
	_ *tls.Config,
	_ bool,
	_ *utils.RTTStats,
	_ qlogwriter.Recorder,
	_ utils.Logger,
	version protocol.Version,
) CryptoSetup {
	initialSealer, initialOpener := NewInitialAEAD(connID, protocol.PerspectiveServer, version)
	return &cryptoSetup{initialOpener: initialOpener, initialSealer: initialSealer}
}

func (h *cryptoSetup) StartHandshake(context.Context) error {
	return errTinyGoQUICUnsupported
}

func (h *cryptoSetup) Close() error {
	return nil
}

func (h *cryptoSetup) ChangeConnectionID(id protocol.ConnectionID) {
	initialSealer, initialOpener := NewInitialAEAD(id, protocol.PerspectiveClient, protocol.Version1)
	h.initialSealer = initialSealer
	h.initialOpener = initialOpener
}

func (h *cryptoSetup) GetSessionTicket() ([]byte, error) {
	return nil, errTinyGoQUICUnsupported
}

func (h *cryptoSetup) HandleMessage([]byte, protocol.EncryptionLevel) error {
	return errTinyGoQUICUnsupported
}

func (h *cryptoSetup) NextEvent() Event {
	return Event{Kind: EventNoEvent}
}

func (h *cryptoSetup) SetLargest1RTTAcked(protocol.PacketNumber) error {
	return errTinyGoQUICUnsupported
}

func (h *cryptoSetup) DiscardInitialKeys() {}

func (h *cryptoSetup) SetHandshakeConfirmed() {}

func (h *cryptoSetup) ConnectionState() ConnectionState {
	return ConnectionState{}
}

func (h *cryptoSetup) GetInitialOpener() (LongHeaderOpener, error) {
	if h.initialOpener == nil {
		return nil, ErrKeysDropped
	}
	return h.initialOpener, nil
}

func (h *cryptoSetup) GetHandshakeOpener() (LongHeaderOpener, error) {
	return nil, ErrKeysNotYetAvailable
}

func (h *cryptoSetup) Get0RTTOpener() (LongHeaderOpener, error) {
	return nil, ErrKeysNotYetAvailable
}

func (h *cryptoSetup) Get1RTTOpener() (ShortHeaderOpener, error) {
	return nil, ErrKeysNotYetAvailable
}

func (h *cryptoSetup) GetInitialSealer() (LongHeaderSealer, error) {
	if h.initialSealer == nil {
		return nil, ErrKeysDropped
	}
	return h.initialSealer, nil
}

func (h *cryptoSetup) GetHandshakeSealer() (LongHeaderSealer, error) {
	return nil, ErrKeysNotYetAvailable
}

func (h *cryptoSetup) Get0RTTSealer() (LongHeaderSealer, error) {
	return nil, ErrKeysNotYetAvailable
}

func (h *cryptoSetup) Get1RTTSealer() (ShortHeaderSealer, error) {
	return nil, ErrKeysNotYetAvailable
}
