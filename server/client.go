package Server

import (
	"errors"
	"fmt"
	"github.com/mastercactapus/go-minecraft-gateway/packet-decoder"
	"github.com/mastercactapus/go-minecraft-gateway/packet-encoder"
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
	State         int
	Conn          net.Conn
	UUID          string
	Authenticated bool
	Username      string
	Reader        io.Reader
	Writer        io.Writer
	Decoder       *PacketDecoder.Decoder
	Encoder       *PacketEncoder.Encoder
}

func (self Server) NewClient(conn net.Conn) {
	c := new(ClientConnection)
	c.Conn = conn

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Client terminated: %s -- %s\n", c.Conn.RemoteAddr().String(), err)
		}
		c.Conn.Close()
	}()

	c.Reader = conn
	c.Writer = conn

	c.Decoder = PacketDecoder.NewDecoder(c.Reader)
	c.Encoder = PacketEncoder.NewEncoder(c.Writer)

	nextState := c.DoHandshake()

	if nextState == STATUS {
		c.DoStatusCheck()
		c.DoStatusPing()
	} else if nextState == LOGIN {
		self.AuthenticateClient(c)
	} else {
		panic("Invalid state from client")
	}
}
