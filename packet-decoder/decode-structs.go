package PacketDecoder

import (
	"github.com/mastercactapus/go-minecraft-gateway/packets"
)

func (self *Decoder) Handshake() *Packets.Handshake {
	packet := new(Packets.Handshake)
	packet.ID = 0
	packet.ProtocolVersion = self.Varint()
	packet.ServerAddress = self.String()
	packet.ServerPort = self.Ushort()
	packet.NextState = self.Varint()

	return packet
}
func (self *Decoder) StatusPing() *Packets.StatusPing {
	packet := new(Packets.StatusPing)
	packet.ID = 1
	packet.Time = self.Long()
	return packet
}

func (self *Decoder) LoginStart() *Packets.LoginStart {
	packet := new(Packets.LoginStart)
	packet.ID = 0
	packet.Name = self.String()
	return packet
}

func (self *Decoder) EncryptionResponse() *Packets.EncryptionResponse {
	packet := new(Packets.EncryptionResponse)
	packet.ID = 1
	secretLength := self.Short()
	packet.SharedSecret = self.readBytesPanic(uint64(secretLength))
	verifyLength := self.Short()
	packet.VerifyToken = self.readBytesPanic(uint64(verifyLength))
	return packet
}
