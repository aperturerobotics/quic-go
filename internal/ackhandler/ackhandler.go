package ackhandler

import (
	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/utils"
)

// NewAckHandler creates a new SentPacketHandler and a new ReceivedPacketHandler
func NewAckHandler(
	initialPacketNumber protocol.PacketNumber,
	rttStats *utils.RTTStats,
	pers protocol.Perspective,
	logger utils.Logger,
	version protocol.VersionNumber,
) (SentPacketHandler, ReceivedPacketHandler) {
	sph := newSentPacketHandler(initialPacketNumber, rttStats, pers, logger)
	return sph, newReceivedPacketHandler(sph, rttStats, logger, version)
}
