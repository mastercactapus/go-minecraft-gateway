package PacketDecoder

import (
	"github.com/mastercactapus/go-minecraft-gateway/packets-serverbound"
)

func (self *Decoder) ServerboundHandshake() *Serverbound.Handshake {
	packet := new(Serverbound.Handshake)
	packet.ID = 0
	packet.ProtocolVersion = self.Varint()
	packet.ServerAddress = self.String()
	packet.ServerPort = self.Ushort()
	packet.NextState = self.Varint()

	return packet
}
func (self *Decoder) ServerboundStatusPing() *Serverbound.StatusPing {
	packet := new(Serverbound.StatusPing)
	packet.ID = 1
	packet.Time = self.Long()
	return packet
}

func (self *Decoder) ServerboundLoginStart() *Serverbound.LoginStart {
	packet := new(Serverbound.LoginStart)
	packet.ID = 0
	packet.Name = self.String()
	return packet
}

func (self *Decoder) ServerboundEncryptionResponse() *Serverbound.EncryptionResponse {
	packet := new(Serverbound.EncryptionResponse)
	packet.ID = 1
	secretLength := self.Short()
	packet.SharedSecret = self.readBytesPanic(uint64(secretLength))
	verifyLength := self.Short()
	packet.VerifyToken = self.readBytesPanic(uint64(verifyLength))
	return packet
}
