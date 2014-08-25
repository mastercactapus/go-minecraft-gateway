package Server

import (
	"crypto/cipher"
	"errors"
	"fmt"
	"github.com/mastercactapus/go-minecraft-gateway/packet-decoder"
	"github.com/mastercactapus/go-minecraft-gateway/packet-encoder"
	"net"
	"runtime/debug"
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
	ServerHash    string
	Texture       string
	Username      string
	Cipher        cipher.Block
	Decoder       *PacketDecoder.Decoder
	Encoder       *PacketEncoder.Encoder
}

func (self Server) NewClient(conn net.Conn) {
	c := new(ClientConnection)
	c.Conn = conn

	defer func() {
		err := recover()
		if c.Authenticated {
			self.OnlineClients--
		}
		if err != nil {
			fmt.Printf("Client terminated: %s -- %s\n", c.Conn.RemoteAddr().String(), err)
			debug.PrintStack()
		}
		c.Conn.Close()
	}()

	c.Decoder = PacketDecoder.NewDecoder(c.Conn)
	c.Encoder = PacketEncoder.NewEncoder(c.Conn)

	nextState := self.DoHandshake(c)

	if nextState == STATUS {
		self.DoStatusCheck(c)
		self.DoStatusPing(c)
	} else if nextState == LOGIN {
		self.AuthenticateClient(c)
		fmt.Printf("User %s has joined the server. id=%s \n", c.Username, c.UUID)
	} else {
		panic("Invalid state from client")
	}
}
