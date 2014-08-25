package Server

import (
	"fmt"
	"github.com/mastercactapus/go-minecraft-gateway/packets"
)

func (self *Server) AuthenticateClient(client *ClientConnection) {
	packet := client.Decoder.Packet()

	packetType := packet.Varint()

	if packetType != 0 {
		panic(UnexpectedPacketType)
	}

	loginStart := packet.LoginStart()
	fmt.Println("Login:", loginStart.Name)

	encryptionRequest := new(Packets.EncryptionRequest)

}
