package Server

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"net"
)

const (
	MaxClientsDefault = 1024
)

type Server struct {
	Connections   []ClientConnection
	PrivateKey    *rsa.PrivateKey
	PublicKeyData []byte
	Socket        net.Listener
	ServerID      string
	MaxClients    int
	OnlineClients int
}

func HandleConnections(s *Server) {
	for {
		conn, err := s.Socket.Accept()
		if err != nil {
			s.Socket.Close()
			break
		}

		s.NewClient(conn)
	}

}

func NewServer(bindAddress string) (*Server, error) {
	s := new(Server)
	sck, err := net.Listen("tcp", bindAddress)

	//seems to be blank nowadays
	s.ServerID = ""
	s.MaxClients = MaxClientsDefault
	s.OnlineClients = 0

	if err != nil {
		panic(err)
	}

	s.Socket = sck
	s.PrivateKey, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	s.PublicKeyData, err = x509.MarshalPKIXPublicKey(&s.PrivateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	HandleConnections(s)

	return s, nil
}
