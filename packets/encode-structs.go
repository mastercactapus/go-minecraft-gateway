package Packets

import (
	"encoding/json"
)

func (s *Stream) WriteStatusResponse(data *StatusResponse) {
	s.writeUvarint(data.ID)
	jsonString, err := json.Marshal(data.JSONResponse)
	if err != nil {
		panic(err)
	}
	s.writeUvarint(uint64(len(jsonString)))
	s.writeBytes(jsonString)
	s.writePacket()
}

func (s *Stream) WriteStatusPong(data *StatusPong) {
	s.writeUvarint(data.ID)
	s.writeLong(data.Time)
	s.writePacket()
}

func (s *Stream) Disconnect(data *Disconnect) {
	s.writeUvarint(data.ID)
	s.writeString(data.JSONData)
	s.writePacket()
}
func (s *Stream) EncryptionRequest(data *EncryptionRequest) {
	s.writeUvarint(data.ID)
	s.writeString(data.ServerID)
	s.writeShort(int16(len(data.PublicKey)))
	s.writeBytes(data.PublicKey)
	s.writeShort(int16(len(data.VerifyToken)))
	s.writeBytes(data.VerifyToken)
	s.writePacket()
}
func (s *Stream) LoginSuccess(data *LoginSuccess) {
	s.writeUvarint(data.ID)
	s.writeString(data.UUID)
	s.writeString(data.Username)
	s.writePacket()
}

func (s *Stream) Handshake(data *Handshake) {
	s.writeUvarint(data.ID)
	s.writeUvarint(data.ProtocolVersion)
	s.writeString(data.ServerAddress)
	s.writeUshort(data.ServerPort)
	s.writeUvarint(data.NextState)
}

func (s *Stream) LoginStart(data *LoginStart) {
	s.writeUvarint(data.ID)
	s.writeString(data.Name)
	s.writePacket()
}
