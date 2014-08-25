package Server

import (
	"errors"
	"fmt"
	"github.com/mastercactapus/go-minecraft-gateway/packet-decoder"
	"github.com/mastercactapus/go-minecraft-gateway/packet-encoder"
	"github.com/mastercactapus/go-minecraft-gateway/packets-clientbound"
	"io"
	"net"
)

const (
	HANDSHAKE = 0
	STATUS    = 1
	LOGIN     = 2
	PLAY      = 3

	ProtocolVersion = 5
	VersionName     = "1.7.10"
)

var UnexpectedPacketType = errors.New("unexpected packet type for current state")
var UnsupportedProtocolVersion = errors.New("unsupported protocol version from client")

type ClientConnection struct {
	State   int
	Conn    net.Conn
	UUID    string
	Reader  io.Reader
	Writer  io.Writer
	Decoder *PacketDecoder.Decoder
	Encoder *PacketEncoder.Encoder
}

func (self Server) NewClient(conn net.Conn) {
	c := new(ClientConnection)
	c.Conn = conn

	c.Reader = conn
	c.Writer = conn

	c.Decoder = PacketDecoder.NewDecoder(c.Reader)
	c.Encoder = PacketEncoder.NewEncoder(c.Writer)

	nextState, err := c.DoHandshake()
	if err != nil {
		fmt.Printf("Handshake failed: %s\n", err)
		c.Conn.Close()
		return
	}
	if nextState == STATUS {
		err = c.DoStatus()
		if err != nil {
			fmt.Printf("Status check failed: %s\n", err)
			c.Conn.Close()
			return
		}
	} else if nextState == LOGIN {
		fmt.Println("Login not implemented")
		c.Conn.Close()
		return
	} else {
		fmt.Println("Invalid state from client")
		c.Conn.Close()
		return
	}

	c.Conn.Close()
}

func (self ClientConnection) DoStatus() error {
	packet, err := self.Decoder.Packet()
	if err != nil {
		return err
	}

	packetType, err := packet.Varint()
	if err != nil {
		return err
	}
	if packetType != 0 {
		return UnexpectedPacketType
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
	err = response.ClientboundStatusResponse(data)
	if err != nil {
		return err
	}

	err = self.Encoder.Packet(response)
	if err != nil {
		return err
	}

	return self.DoStatusPing()
}

func (self ClientConnection) DoStatusPing() error {
	packet, err := self.Decoder.Packet()
	if err != nil {
		return err
	}

	packetType, err := packet.Varint()
	if err != nil {
		return err
	}
	if packetType != 1 {
		return UnexpectedPacketType
	}

	data, err := packet.ServerboundStatusPing()
	if err != nil {
		return err
	}

	response := PacketEncoder.NewPacket()
	err = response.ClientboundStatusPing(&Clientbound.StatusPing{1, data.Time})
	if err != nil {
		return err
	}

	return self.Encoder.Packet(response)

}

func (self ClientConnection) DoHandshake() (uint64, error) {
	packet, err := self.Decoder.Packet()
	if err != nil {
		return 0, err
	}
	packetType, err := packet.Varint()
	if err != nil {
		return 0, err
	}
	if packetType != 0 {
		return 0, UnexpectedPacketType
	}

	handshake, err := packet.ServerboundHandshake()
	if err != nil {
		return 0, err
	}
	if handshake.ProtocolVersion != ProtocolVersion {
		return 0, UnsupportedProtocolVersion
	}

	return handshake.NextState, nil
}
