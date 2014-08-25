package Server

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"github.com/mastercactapus/go-minecraft-gateway/auth"
	"github.com/mastercactapus/go-minecraft-gateway/auth/cfb8"
	"github.com/mastercactapus/go-minecraft-gateway/packet-decoder"
	"github.com/mastercactapus/go-minecraft-gateway/packet-encoder"
	"github.com/mastercactapus/go-minecraft-gateway/packets"
	"io"
)

func (self *Server) AuthenticateClient(client *ClientConnection) {
	packet := client.Decoder.Packet()
	packetType := packet.Varint()

	if packetType != 0 {
		panic(UnexpectedPacketType)
	}

	loginStart := packet.LoginStart()
	client.Username = loginStart.Name

	encryptionRequest := new(Packets.EncryptionRequest)
	encryptionRequest.ID = 1
	encryptionRequest.ServerID = self.ServerID
	encryptionRequest.PublicKey = self.PublicKeyData
	encryptionRequest.VerifyToken = make([]byte, 4)
	n, err := rand.Read(encryptionRequest.VerifyToken)
	if err != nil {
		panic(err)
	}
	if n != 4 {
		panic("not enough entropy for clients")
	}

	requestPacket := PacketEncoder.NewPacket()
	requestPacket.EncryptionRequest(encryptionRequest)

	client.Encoder.Packet(requestPacket)

	nextPacket := client.Decoder.Packet()
	nextPacketType := nextPacket.Varint()
	if nextPacketType != 1 {
		panic(UnexpectedPacketType)
	}

	encryptionResponse := nextPacket.EncryptionResponse()
	verifyToken, err := rsa.DecryptPKCS1v15(rand.Reader, self.PrivateKey, encryptionResponse.VerifyToken)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(verifyToken, encryptionRequest.VerifyToken) {
		panic("VerifyToken mismatch")
	}

	sharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, self.PrivateKey, encryptionResponse.SharedSecret)
	if err != nil {
		panic(err)
	}

	//server hash
	h := sha1.New()
	io.WriteString(h, self.ServerID)
	h.Write(sharedSecret)
	h.Write(encryptionRequest.PublicKey)
	client.ServerHash = Auth.HashFix(h.Sum(nil))

	client.ValidateOnline()

	if err != nil {
		panic(err)
	}
	encryptedStream := CFB8.New(client.Conn, sharedSecret)

	client.Decoder = PacketDecoder.NewDecoder(encryptedStream)
	client.Encoder = PacketEncoder.NewEncoder(encryptedStream)

	loginSuccess := new(Packets.LoginSuccess)
	loginSuccess.ID = 2
	loginSuccess.UUID = client.UUID
	loginSuccess.Username = client.Username

	successPacket := PacketEncoder.NewPacket()
	successPacket.LoginSuccess(loginSuccess)

	client.Encoder.Packet(successPacket)

	client.Authenticated = true
	self.OnlineClients++
}
