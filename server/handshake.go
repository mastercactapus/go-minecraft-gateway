package Server

func (self *Server) DoHandshake(client *ClientConnection) uint64 {
	packet := client.Decoder.Packet()
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
