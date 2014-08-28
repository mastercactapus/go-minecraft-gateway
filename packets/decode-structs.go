package Packets

import (
	"errors"
)

var InvalidPacketID = errors.New("invalid packet id for type")

func (s *Stream) ReadHandshake() *Handshake {
	packet := new(Handshake)
	packet.ID = s.PeekPacketType()
	if packet.ID != 0 {
		panic(InvalidPacketID)
	}
	packet.ProtocolVersion = s.readUvarint()
	packet.ServerAddress = s.readString()
	packet.ServerPort = s.readUshort()
	packet.NextState = s.readUvarint()
	s.hasPacket = false

	return packet
}

func (s *Stream) ReadStatusRequest() *StatusRequest {
	packet := new(StatusRequest)
	packet.ID = s.PeekPacketType()
	if packet.ID != 0 {
		panic(InvalidPacketID)
	}
	s.hasPacket = false
	return packet
}

func (s *Stream) ReadStatusPing() *StatusPing {
	packet := new(StatusPing)
	packet.ID = s.PeekPacketType()
	if packet.ID != 1 {
		panic(InvalidPacketID)
	}
	packet.Time = s.readLong()
	s.hasPacket = false
	return packet
}

func (s *Stream) ReadLoginStart() *LoginStart {
	packet := new(LoginStart)
	packet.ID = s.PeekPacketType()
	if packet.ID != 0 {
		panic(InvalidPacketID)
	}
	packet.Name = s.readString()
	s.hasPacket = false
	return packet
}

func (s *Stream) ReadEncryptionResponse() *EncryptionResponse {
	packet := new(EncryptionResponse)
	packet.ID = s.PeekPacketType()
	if packet.ID != 1 {
		panic(InvalidPacketID)
	}
	secretLength := s.readShort()
	packet.SharedSecret = s.readBytes(uint64(secretLength))
	verifyLength := s.readShort()
	packet.VerifyToken = s.readBytes(uint64(verifyLength))
	s.hasPacket = false
	return packet
}

func (s *Stream) ReadLoginSuccess() *LoginSuccess {
	packet := new(LoginSuccess)
	packet.ID = s.PeekPacketType()
	if packet.ID != 2 {
		panic(InvalidPacketID)
	}
	packet.UUID = s.readString()
	packet.Username = s.readString()
	s.hasPacket = false
	return packet
}
