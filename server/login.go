package Server

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"github.com/mastercactapus/go-minecraft-gateway/auth"
	"github.com/mastercactapus/go-minecraft-gateway/auth/cfb8"
	"github.com/mastercactapus/go-minecraft-gateway/packets"
	"io"
)

func (s *Server) AuthenticateClient(c *ClientConnection) {
	loginStart := c.packetStream.ReadLoginStart()

	c.Username = loginStart.Name
	encryptionRequest := new(Packets.EncryptionRequest)
	encryptionRequest.ID = 1
	encryptionRequest.ServerID = s.ServerID
	encryptionRequest.PublicKey = s.PublicKeyData
	encryptionRequest.VerifyToken = make([]byte, 4)
	n, err := rand.Read(encryptionRequest.VerifyToken)
	if err != nil {
		panic(err)
	}
	if n != 4 {
		panic("not enough entropy for clients")
	}

	c.packetStream.WriteEncryptionRequest(encryptionRequest)
	encryptionResponse := c.packetStream.ReadEncryptionResponse()

	verifyToken, err := rsa.DecryptPKCS1v15(rand.Reader, s.PrivateKey, encryptionResponse.VerifyToken)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(verifyToken, encryptionRequest.VerifyToken) {
		panic("VerifyToken mismatch")
	}

	sharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, s.PrivateKey, encryptionResponse.SharedSecret)
	if err != nil {
		panic(err)
	}

	//server hash
	h := sha1.New()
	io.WriteString(h, s.ServerID)
	h.Write(sharedSecret)
	h.Write(encryptionRequest.PublicKey)
	c.ServerHash = Auth.HashFix(h.Sum(nil))

	c.ValidateOnline()

	if err != nil {
		panic(err)
	}
	encryptedStream := CFB8.New(c.Conn, sharedSecret)

	c.packetStream = Packets.NewStream(encryptedStream)

	loginSuccess := new(Packets.LoginSuccess)
	loginSuccess.ID = 2
	loginSuccess.UUID = c.UUID
	loginSuccess.Username = c.Username

	c.packetStream.WriteLoginSuccess(loginSuccess)

	c.Authenticated = true
	s.OnlineClients++
}
