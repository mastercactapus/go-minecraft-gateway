package PacketDecoder

import "../packets-serverbound"

func (self *Decoder) ServerboundHandshake() (*Serverbound.Handshake, error) {
	var err error
	packet := new(Serverbound.Handshake)
	packet.ID = 0
	packet.ProtocolVersion, err = self.Varint()
	if err != nil {
		return packet, err
	}
	packet.ServerAddress, err = self.String()
	if err != nil {
		return packet, err
	}
	packet.ServerPort, err = self.Ushort()
	if err != nil {
		return packet, err
	}
	packet.NextState, err = self.Varint()
	if err != nil {
		return packet, err
	}

	return packet, nil
}
func (self *Decoder) ServerboundStatusPing() (*Serverbound.StatusPing, error) {
	var err error
	packet := new(Serverbound.StatusPing)
	packet.ID = 1
	packet.Time, err = self.Long()
	return packet, err
}
