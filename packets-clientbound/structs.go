package Clientbound

//Status State
type StatusResponse struct {
	ID           uint64 //0x00
	JSONResponse struct {
		Version struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		} `json:"version"`
		Players struct {
			Max    int `json:"max"`
			Online int `json:"online"`
		} `json:"players"`
		Description string `json:"description"`
	}
}
type StatusPing struct {
	ID   uint64 //0x01
	Time int64
}

//Login State
type Disconnect struct {
	ID       uint64 //0x00
	JSONData string
}
type EncryptionRequest struct {
	ID             uint64 //0x01
	ServerID       string
	PublicKeyLen   int16
	PublicKey      []byte
	VerifyTokenLen int16
	VerifyToken    []byte
}
type LoginSuccess struct {
	ID       uint64 //0x02
	UUID     string
	Username string
}
