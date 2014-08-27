package Packets

func (s *Stream) ReadHandshake() *Handshake {
	packet := new(Handshake)
	packet.ID = 0
	packet.ProtocolVersion = s.readUvarint()
	packet.ServerAddress = s.readString()
	packet.ServerPort = s.readUshort()
	packet.NextState = s.readUvarint()

	return packet
}
func (s *Stream) ReadStatusPing() *StatusPing {
	packet := new(StatusPing)
	packet.ID = 1
	packet.Time = s.readLong()
	return packet
}

func (s *Stream) ReadLoginStart() *LoginStart {
	packet := new(LoginStart)
	packet.ID = 0
	packet.Name = s.readString()
	return packet
}

func (s *Stream) ReadEncryptionResponse() *EncryptionResponse {
	packet := new(EncryptionResponse)
	packet.ID = 1
	secretLength := s.readShort()
	packet.SharedSecret = s.readBytes(uint64(secretLength))
	verifyLength := s.readShort()
	packet.VerifyToken = s.readBytes(uint64(verifyLength))
	return packet
}

func (s *Stream) ReadLoginSuccess() *LoginSuccess {
	packet := new(LoginSuccess)
	packet.ID = 2
	packet.UUID = s.readString()
	packet.Username = s.readString()
	return packet
}
