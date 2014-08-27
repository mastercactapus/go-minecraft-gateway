package Server

func (s *Server) DoHandshake(c *ClientConnection) uint64 {
	handshake := c.packetStream.ReadHandshake()

	if handshake.ProtocolVersion != ProtocolVersion {
		panic(UnsupportedProtocolVersion)
	}

	return handshake.NextState
}
