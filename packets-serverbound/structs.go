package Serverbound

//Handshake State
type Handshake struct {
	ID              uint64 //0x00
	ProtocolVersion uint64
	ServerAddress   string
	ServerPort      uint16
	NextState       uint64
}

//Status State
type StatusRequest struct {
	ID uint64 //0x00
}
type StatusPing struct {
	ID   uint64 //0x01
	Time int64
}

//Login State
type LoginStart struct {
	ID   uint64 //ID 0x00
	Name string
}
type EncryptionResponse struct {
	ID              uint64 //0x01
	SharedSecretLen int16
	SharedSecret    []byte
	VerifyTokenLen  int16
	VerifyToken     []byte
}
