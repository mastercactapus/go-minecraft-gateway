package PacketEncoder

import (
	"encoding/json"
	"github.com/mastercactapus/go-minecraft-gateway/packets-clientbound"
)

func (self *Encoder) ClientboundStatusResponse(data *Clientbound.StatusResponse) error {
	err := self.Varint(data.ID)
	if err != nil {
		return err
	}

	jsonString, err := json.Marshal(data.JSONResponse)
	if err != nil {
		return err
	}

	err = self.Varint(uint64(len(jsonString)))
	if err != nil {
		return err
	}

	return self.WriteBytes(jsonString)
}

func (self *Encoder) ClientboundStatusPing(data *Clientbound.StatusPing) error {
	err := self.Varint(data.ID)
	if err != nil {
		return err
	}
	return self.Long(data.Time)
}
