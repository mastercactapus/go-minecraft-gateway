package Server

import (
	"github.com/mastercactapus/go-minecraft-gateway/packet-encoder"
	"github.com/mastercactapus/go-minecraft-gateway/packets-clientbound"
)

func (self ClientConnection) DoStatusCheck() {
	packet := self.Decoder.Packet()
	packetType := packet.Varint()

	if packetType != 0 {
		panic(UnexpectedPacketType)
	}

	//create response packet
	data := new(Clientbound.StatusResponse)
	data.ID = 0

	data.JSONResponse.Version.Name = VersionName
	data.JSONResponse.Version.Protocol = ProtocolVersion

	data.JSONResponse.Players.Max = 9001
	data.JSONResponse.Players.Online = 2
	data.JSONResponse.Description = "Minecraft Gateway Server"

	response := PacketEncoder.NewPacket()
	response.ClientboundStatusResponse(data)
	self.Encoder.Packet(response)

}

func (self ClientConnection) DoStatusPing() {
	packet := self.Decoder.Packet()
	packetType := packet.Varint()

	if packetType != 1 {
		panic(UnexpectedPacketType)
	}

	data := packet.ServerboundStatusPing()
	response := PacketEncoder.NewPacket()
	response.ClientboundStatusPing(&Clientbound.StatusPing{1, data.Time})

	self.Encoder.Packet(response)

}