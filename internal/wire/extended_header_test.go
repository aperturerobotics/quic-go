package wire

import (
	"testing"

	"github.com/quic-go/quic-go/internal/protocol"
	"github.com/quic-go/quic-go/quicvarint"

	"github.com/stretchr/testify/require"
)

func TestWritesLongHeaderVersion1(t *testing.T) {
	header := &ExtendedHeader{
		Header: Header{
			Type:             protocol.PacketTypeHandshake,
			DestConnectionID: protocol.ParseConnectionID([]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe}),
			SrcConnectionID:  protocol.ParseConnectionID([]byte{0xde, 0xca, 0xfb, 0xad, 0x0, 0x0, 0x13, 0x37}),
			Version:          0x1020304,
			Length:           1234,
		},
		PacketNumber:    0xdecaf,
		PacketNumberLen: protocol.PacketNumberLen3,
	}
	b, err := header.Append(nil, protocol.Version1)
	require.NoError(t, err)
	expected := []byte{
		0xc0 | 0x2<<4 | 0x2,
		0x1, 0x2, 0x3, 0x4, // version number
		0x6,                                // dest connection ID length
		0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, // dest connection ID
		0x8,                                          // src connection ID length
		0xde, 0xca, 0xfb, 0xad, 0x0, 0x0, 0x13, 0x37, // source connection ID
	}
	expected = append(expected, encodeVarInt(1234)...)      // length
	expected = append(expected, []byte{0xd, 0xec, 0xaf}...) // packet number
	require.Equal(t, expected, b)
	require.Equal(t, protocol.ByteCount(len(b)), header.GetLength(protocol.Version1))
}

func TestWritesHandshakePacketVersion2(t *testing.T) {
	header := &ExtendedHeader{
		Header: Header{
			Version: protocol.Version2,
			Type:    protocol.PacketTypeHandshake,
		},
		PacketNumber:    0xdecafbad,
		PacketNumberLen: protocol.PacketNumberLen4,
	}
	b, err := header.Append(nil, protocol.Version2)
	require.NoError(t, err)
	require.Equal(t, byte(0b11), b[0]>>4&0b11)
	require.Equal(t, protocol.ByteCount(len(b)), header.GetLength(protocol.Version2))
}

func TestWritesHeaderWith20ByteConnectionID(t *testing.T) {
	srcConnID := protocol.ParseConnectionID([]byte{1, 2, 3, 4, 5, 6, 7, 8})
	header := &ExtendedHeader{
		Header: Header{
			SrcConnectionID:  srcConnID,
			DestConnectionID: protocol.ParseConnectionID([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}), // connection IDs must be at most 20 bytes long
			Version:          0x1020304,
			Type:             0x5,
		},
		PacketNumber:    0xdecafbad,
		PacketNumberLen: protocol.PacketNumberLen4,
	}
	b, err := header.Append(nil, protocol.Version1)
	require.NoError(t, err)
	require.Contains(t, string(b), string([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}))
	require.Equal(t, protocol.ByteCount(len(b)), header.GetLength(protocol.Version1))
}

func TestWritesInitialContainingToken(t *testing.T) {
	token := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")
	header := &ExtendedHeader{
		Header: Header{
			Version: 0x1020304,
			Type:    protocol.PacketTypeInitial,
			Token:   token,
		},
		PacketNumber:    0xdecafbad,
		PacketNumberLen: protocol.PacketNumberLen4,
	}
	b, err := header.Append(nil, protocol.Version1)
	require.NoError(t, err)
	require.Equal(t, byte(0), b[0]>>4&0b11)
	expectedSubstring := append(encodeVarInt(uint64(len(token))), token...)
	require.Contains(t, string(b), string(expectedSubstring))
	require.Equal(t, protocol.ByteCount(len(b)), header.GetLength(protocol.Version1))
}

func TestUses2ByteEncodingForLengthOnInitialPackets(t *testing.T) {
	header := &ExtendedHeader{
		Header: Header{
			Version: 0x1020304,
			Type:    protocol.PacketTypeInitial,
			Length:  37,
		},
		PacketNumber:    0xdecafbad,
		PacketNumberLen: protocol.PacketNumberLen4,
	}
	b, err := header.Append(nil, protocol.Version1)
	require.NoError(t, err)
	lengthEncoded := quicvarint.AppendWithLen(nil, 37, 2)
	require.Equal(t, lengthEncoded, b[len(b)-6:len(b)-4])
	require.Equal(t, protocol.ByteCount(len(b)), header.GetLength(protocol.Version1))
}

func TestWritesInitialPacketVersion2(t *testing.T) {
	header := &ExtendedHeader{
		Header: Header{
			Version: protocol.Version2,
			Type:    protocol.PacketTypeInitial,
		},
		PacketNumber:    0xdecafbad,
		PacketNumberLen: protocol.PacketNumberLen4,
	}
	b, err := header.Append(nil, protocol.Version2)
	require.NoError(t, err)
	require.Equal(t, byte(0b01), b[0]>>4&0b11)
	require.Equal(t, protocol.ByteCount(len(b)), header.GetLength(protocol.Version2))
}

func TestWrites0RTTPacketVersion2(t *testing.T) {
	header := &ExtendedHeader{
		Header: Header{
			Version: protocol.Version2,
			Type:    protocol.PacketType0RTT,
		},
		PacketNumber:    0xdecafbad,
		PacketNumberLen: protocol.PacketNumberLen4,
	}
	b, err := header.Append(nil, protocol.Version2)
	require.NoError(t, err)
	require.Equal(t, byte(0b10), b[0]>>4&0b11)
	require.Equal(t, protocol.ByteCount(len(b)), header.GetLength(protocol.Version2))
}

func TestWritesRetryPacket(t *testing.T) {
	token := []byte("Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.")

	for _, version := range []protocol.Version{protocol.Version1, protocol.Version2} {
		t.Run(version.String(), func(t *testing.T) {
			header := &ExtendedHeader{Header: Header{
				Version: version,
				Type:    protocol.PacketTypeRetry,
				Token:   token,
			}}
			b, err := header.Append(nil, version)
			require.NoError(t, err)

			var expected []byte
			switch version {
			case protocol.Version1:
				expected = append(expected, 0xc0|0b11<<4)
			case protocol.Version2:
				expected = append(expected, 0xc0)
			}

			expected = appendVersion(expected, version)
			expected = append(expected, 0x0) // dest connection ID length
			expected = append(expected, 0x0) // src connection ID length
			expected = append(expected, token...)
			require.Equal(t, expected, b)
		})
	}
}

func BenchmarkParseExtendedHeader(b *testing.B) {
	data, err := (&ExtendedHeader{
		Header: Header{
			Type:             protocol.PacketTypeHandshake,
			DestConnectionID: protocol.ParseConnectionID([]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe}),
			SrcConnectionID:  protocol.ParseConnectionID([]byte{0xde, 0xca, 0xfb, 0xad, 0x0, 0x0, 0x13, 0x37}),
			Version:          protocol.Version1,
			Length:           1234,
		},
		PacketNumber:    0xdecaf,
		PacketNumberLen: protocol.PacketNumberLen3,
	}).Append(nil, protocol.Version1)
	if err != nil {
		b.Fatal(err)
	}
	data = append(data, make([]byte, 1231)...)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		hdr, _, _, err := ParsePacket(data)
		if err != nil {
			b.Fatal(err)
		}
		if _, err := hdr.ParseExtended(data); err != nil {
			b.Fatal(err)
		}
	}
}
