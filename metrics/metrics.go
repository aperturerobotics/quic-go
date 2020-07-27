package metrics

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// Measures
var (
	connections = stats.Int64("quic-go/connections", "number of QUIC connections", stats.UnitDimensionless)
	lostPackets = stats.Int64("quic-go/lost-packets", "number of packets declared lost", stats.UnitDimensionless)
	sentPackets = stats.Int64("quic-go/sent-packets", "number of packets sent", stats.UnitDimensionless)
	ptos        = stats.Int64("quic-go/ptos", "number of times the PTO timer fired", stats.UnitDimensionless)
	closes      = stats.Int64("quic-go/close", "number of connections closed", stats.UnitDimensionless)
)

// Tags
var (
	keyPerspective, _      = tag.NewKey("perspective")
	keyIPVersion, _        = tag.NewKey("ip_version")
	keyEncryptionLevel, _  = tag.NewKey("encryption_level")
	keyPacketLossReason, _ = tag.NewKey("packet_loss_reason")
	keyPacketType, _       = tag.NewKey("packet_type")
	keyCloseReason, _      = tag.NewKey("close_reason")
	keyCloseRemote, _      = tag.NewKey("close_remote")
	keyErrorCode, _        = tag.NewKey("error_code")
	keyHandshakePhase, _   = tag.NewKey("handshake_phase")
)

// Views
var (
	ConnectionsView = &view.View{
		Measure:     connections,
		TagKeys:     []tag.Key{keyPerspective, keyIPVersion},
		Aggregation: view.Count(),
	}
	LostPacketsView = &view.View{
		Measure:     lostPackets,
		TagKeys:     []tag.Key{keyEncryptionLevel, keyPacketLossReason},
		Aggregation: view.Count(),
	}
	SentPacketsView = &view.View{
		Measure:     sentPackets,
		TagKeys:     []tag.Key{keyPacketType},
		Aggregation: view.Count(),
	}
	PTOView = &view.View{
		Measure:     ptos,
		TagKeys:     []tag.Key{keyHandshakePhase},
		Aggregation: view.Count(),
	}
	CloseView = &view.View{
		Measure:     closes,
		TagKeys:     []tag.Key{keyCloseReason, keyErrorCode},
		Aggregation: view.Count(),
	}
)

// DefaultViews collects all OpenCensus views for metric gathering purposes
var DefaultViews = []*view.View{
	ConnectionsView,
	LostPacketsView,
	SentPacketsView,
	CloseView,
}
