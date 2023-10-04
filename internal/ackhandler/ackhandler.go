package ackhandler

import (
	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/internal/utils"
)

// NewAckHandler creates a new SentPacketHandler and a new ReceivedPacketHandler.
// clientAddressValidated indicates whether the address was validated beforehand by an address validation token.
// clientAddressValidated has no effect for a client.
func NewAckHandler(
	initialPacketNumber protocol.PacketNumber,
	initialMaxDatagramSize protocol.ByteCount,
	rttStats *utils.RTTStats,
	clientAddressValidated bool,
	enableECN bool,
	pers protocol.Perspective,
	logger utils.Logger,
) (SentPacketHandler, ReceivedPacketHandler) {
	sph := newSentPacketHandler(initialPacketNumber, initialMaxDatagramSize, rttStats, clientAddressValidated, enableECN, pers, logger)
	return sph, newReceivedPacketHandler(sph, rttStats, logger)
}
