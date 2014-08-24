package Server

import (
	"../packet-decoder"
	"../packet-encoder"
	"bufio"
	"net"
)

const (
	HANDSHAKE = 0
	STATUS    = 1
	LOGIN     = 2
	PLAY      = 3
)

type ClientConnection struct {
	State   int
	Conn    net.Conn
	UUID    string
	Reader  *bufio.Reader
	Writer  *bufio.Writer
	Decoder *PacketDecoder.Decoder
	Encoder *PacketEncoder.Encoder
}

func (self Server) NewClient(conn net.Conn) {
	c := new(ClientConnection)
	c.Conn = conn

	c.Reader = bufio.NewReader(conn)
	c.Writer = bufio.NewWriter(conn)

	c.Decoder = PacketDecoder.NewDecoder(c.Reader)
	c.Encoder = PacketEncoder.NewEncoder(c.Writer)

}

func (self ClientConnection) DoHandshake() {
	packet, err := self.Decoder.Packet()
}
