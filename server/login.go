package Server

import ()

func (self *Server) AuthenticateClient(client *ClientConnection) {
	packet := client.Decoder.Packet()

	packetType := packet.Varint()

	if packetType != 0 {
		panic(UnexpectedPacketType)
	}

}
