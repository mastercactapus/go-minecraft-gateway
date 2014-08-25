package Server

import (
	"github.com/mastercactapus/go-minecraft-gateway/packet-encoder"
	"github.com/mastercactapus/go-minecraft-gateway/packets"
)

func (self *Server) DoStatusCheck(client *ClientConnection) {
	packet := client.Decoder.Packet()
	packetType := packet.Varint()

	if packetType != 0 {
		panic(UnexpectedPacketType)
	}

	//create response packet
	data := new(Packets.StatusResponse)
	data.ID = 0

	data.JSONResponse.Version.Name = VersionName
	data.JSONResponse.Version.Protocol = ProtocolVersion

	data.JSONResponse.Players.Max = self.MaxClients
	data.JSONResponse.Players.Online = self.OnlineClients
	data.JSONResponse.Description = "Minecraft Gateway Server"

	response := PacketEncoder.NewPacket()
	response.StatusResponse(data)
	client.Encoder.Packet(response)

}

func (self *Server) DoStatusPing(client *ClientConnection) {
	packet := client.Decoder.Packet()
	packetType := packet.Varint()

	if packetType != 1 {
		panic(UnexpectedPacketType)
	}

	data := packet.StatusPing()
	response := PacketEncoder.NewPacket()
	response.StatusPong(&Packets.StatusPong{1, data.Time})

	client.Encoder.Packet(response)

}
