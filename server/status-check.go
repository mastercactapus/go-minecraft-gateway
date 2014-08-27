package Server

import (
	"github.com/mastercactapus/go-minecraft-gateway/packets"
)

func (s *Server) DoStatusCheck(c *ClientConnection) {
	c.packetStream.ReadStatusRequest()

	//create response packet
	statusResponse := new(Packets.StatusResponse)
	statusResponse.ID = 0
	statusResponse.JSONResponse.Version.Name = VersionName
	statusResponse.JSONResponse.Version.Protocol = ProtocolVersion
	statusResponse.JSONResponse.Players.Max = s.MaxClients
	statusResponse.JSONResponse.Players.Online = s.OnlineClients
	statusResponse.JSONResponse.Description = "Minecraft Gateway Server"

	c.packetStream.WriteStatusResponse(statusResponse)
}

func (s *Server) DoStatusPing(c *ClientConnection) {
	statusPing := c.packetStream.ReadStatusPing()
	c.packetStream.WriteStatusPing(statusPing)
}
