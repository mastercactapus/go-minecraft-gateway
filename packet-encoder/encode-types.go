package PacketEncoder

import (
	"encoding/binary"
)

func (self *Encoder) Packet(src *Encoder) error {
	err := self.Varint(src.Written)
	if err != nil {
		return err
	}
	err = self.WriteBytes(src.Data[:src.Written])
	return err
}

func (self *Encoder) Varint(val uint64) error {
	buf := make([]byte, binary.MaxVarintLen32)
	bytes := binary.PutUvarint(buf, val)
	return self.WriteBytes(buf[:bytes])
}

func (self *Encoder) String(str string) error {
	stringLength := uint64(len(str))
	err := self.Varint(stringLength)
	if err != nil {
		return err
	}

	return self.WriteBytes([]byte(str))
}

func (self *Encoder) Bool(val bool) error {
	var b byte
	if val {
		b = 1
	} else {
		b = 0
	}
	return self.Byte(b)
}

func (self *Encoder) Byte(val byte) error {
	return self.WriteByte(val)
}
func (self *Encoder) Short(val int16) error {
	return binary.Write(self, binary.BigEndian, &val)
}
func (self *Encoder) Int(val int32) error {
	return binary.Write(self, binary.BigEndian, &val)
}
func (self *Encoder) Long(val int64) error {
	return binary.Write(self, binary.BigEndian, &val)
}
func (self *Encoder) Float(val float32) error {
	return binary.Write(self, binary.BigEndian, &val)
}
func (self *Encoder) Double(val float64) error {
	return binary.Write(self, binary.BigEndian, &val)
}
