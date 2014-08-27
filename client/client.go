package Client

import (
	"fmt"
	"github.com/mastercactapus/go-minecraft-gateway/packets"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	Conn         net.Conn
	packetStream *Packets.Stream
	Username     string
}

func NewClient(address string, username string) (*Client, error) {
	var err error
	c := new(Client)
	c.Conn, err = net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	c.packetStream = Packets.NewStream(c.Conn)
	c.Username = username

	handshake := new(Packets.Handshake)
	loginStart := new(Packets.LoginStart)

	handshake.ID = 0
	handshake.ProtocolVersion = 5
	spl := strings.Split(address, ":")
	handshake.ServerAddress = spl[0]
	parsedPort, err := strconv.ParseUint(spl[1], 10, 16)
	handshake.ServerPort = uint16(parsedPort)
	if err != nil {
		return nil, err
	}
	handshake.NextState = 2

	loginStart.ID = 0
	loginStart.Name = username

	c.packetStream.WriteHandshake(handshake)
	c.packetStream.WriteLoginStart(loginStart)
	loginSuccess := c.packetStream.ReadLoginSuccess()

	fmt.Printf("Logged in as %s %s\n", loginSuccess.Username, loginSuccess.UUID)

	return c, err
}
