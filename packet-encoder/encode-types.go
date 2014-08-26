package PacketEncoder

import (
	"encoding/binary"
)

func (self *Encoder) Packet(src *Encoder) {
	data := src.bytesBuffer.Bytes()
	self.Varint(uint64(len(data)))
	self.writeBytesPanic(data)
}

func (self *Encoder) Varint(val uint64) {
	buf := make([]byte, binary.MaxVarintLen32)
	bytes := binary.PutUvarint(buf, val)
	self.writeBytesPanic(buf[:bytes])
}

func (self *Encoder) String(str string) {
	stringLength := uint64(len(str))
	self.Varint(stringLength)
	self.writeBytesPanic([]byte(str))
}

func (self *Encoder) Bool(val bool) {
	var b byte
	if val {
		b = 1
	} else {
		b = 0
	}
	self.Byte(b)
}

func (self *Encoder) Byte(val byte) {
	self.writeBytePanic(val)
}

func (self *Encoder) Short(val int16) {
	err := binary.Write(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (self *Encoder) Ushort(val uint16) {
	err := binary.Write(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (self *Encoder) Int(val int32) {
	err := binary.Write(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (self *Encoder) Long(val int64) {
	err := binary.Write(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (self *Encoder) Float(val float32) {
	err := binary.Write(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (self *Encoder) Double(val float64) {
	err := binary.Write(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
