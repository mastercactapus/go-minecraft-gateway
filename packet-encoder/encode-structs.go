package PacketEncoder

import (
	"encoding/json"
	"github.com/mastercactapus/go-minecraft-gateway/packets"
)

func (self *Encoder) StatusResponse(data *Packets.StatusResponse) {
	self.Varint(data.ID)
	jsonString, err := json.Marshal(data.JSONResponse)
	if err != nil {
		panic(err)
	}
	self.Varint(uint64(len(jsonString)))
	self.writeBytesPanic(jsonString)
}

func (self *Encoder) StatusPong(data *Packets.StatusPong) {
	self.Varint(data.ID)
	self.Long(data.Time)
}

func (self *Encoder) Disconnect(data *Packets.Disconnect) {
	self.Varint(data.ID)
	self.String(data.JSONData)
}
func (self *Encoder) EncryptionRequest(data *Packets.EncryptionRequest) {
	self.Varint(data.ID)
	self.String(data.ServerID)
	self.Short(int16(len(data.PublicKey)))
	self.writeBytesPanic(data.PublicKey)
	self.Short(int16(len(data.VerifyToken)))
	self.writeBytesPanic(data.VerifyToken)
}
func (self *Encoder) LoginSuccess(data *Packets.LoginSuccess) {
	self.Varint(data.ID)
	self.String(data.UUID)
	self.String(data.Username)
}

func (self *Encoder) Handshake(data *Packets.Handshake) {
	packet := NewPacket()
	packet.Varint(data.ID)
	packet.Varint(data.ProtocolVersion)
	packet.String(data.ServerAddress)
	packet.Ushort(data.ServerPort)
	packet.Varint(data.NextState)
	self.Packet(packet)
}

func (self *Encoder) LoginStart(data *Packets.LoginStart) {
	packet := NewPacket()
	packet.Varint(data.ID)
	packet.String(data.Name)
	self.Packet(packet)
}
