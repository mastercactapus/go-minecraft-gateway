package Server

func (self ClientConnection) DoHandshake() uint64 {
	packet := self.Decoder.Packet()
	packetType := packet.Varint()
	if packetType != 0 {
		panic(UnexpectedPacketType)
	}

	handshake := packet.Handshake()

	if handshake.ProtocolVersion != ProtocolVersion {
		panic(UnsupportedProtocolVersion)
	}

	return handshake.NextState
}
