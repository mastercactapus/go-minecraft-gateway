package PacketEncoder

import (
	"encoding/json"
	"github.com/mastercactapus/go-minecraft-gateway/packets-clientbound"
)

func (self *Encoder) ClientboundStatusResponse(data *Clientbound.StatusResponse) {
	self.Varint(data.ID)
	jsonString, err := json.Marshal(data.JSONResponse)
	if err != nil {
		panic(err)
	}
	self.Varint(uint64(len(jsonString)))
	self.writeBytesPanic(jsonString)
}

func (self *Encoder) ClientboundStatusPing(data *Clientbound.StatusPing) {
	self.Varint(data.ID)
	self.Long(data.Time)
}

func (self *Encoder) ClientboundDisconnect(data *Clientbound.Disconnect) {
	self.Varint(data.ID)
	self.String(data.JSONData)
}
func (self *Encoder) ClientboundEncryptionRequest(data *Clientbound.EncryptionRequest) {
	self.Varint(data.ID)
	self.String(data.ServerID)
	self.Short(int16(len(data.PublicKey)))
	self.writeBytesPanic(data.PublicKey)
	self.Short(int16(len(data.VerifyToken)))
	self.writeBytesPanic(data.VerifyToken)
}
