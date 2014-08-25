package PacketDecoder

import (
	"bytes"
	"encoding/binary"
)

func (self *Decoder) Packet() *Decoder {
	dataSize := self.Varint()
	data := self.readBytesPanic(dataSize)
	reader := bytes.NewReader(data)
	return NewDecoder(reader)
}

func (self *Decoder) Varint() uint64 {
	val, err := binary.ReadUvarint(self)
	if err != nil {
		panic(err)
	}
	return val
}

func (self *Decoder) String() string {
	stringLength := self.Varint()
	stringData := self.readBytesPanic(stringLength)
	return string(stringData)
}

func (self *Decoder) Bool() bool {
	b := self.Byte()
	if b != 0 && b != 1 {
		panic(InvalidBoolean)
	}

	return (b == 1)
}

func (self *Decoder) Byte() byte {
	return self.readBytePanic()
}

func (self *Decoder) Short() int16 {
	var val int16
	err := binary.Read(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (self *Decoder) Ushort() uint16 {
	var val uint16
	err := binary.Read(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (self *Decoder) Int() int32 {
	var val int32
	err := binary.Read(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (self *Decoder) Long() int64 {
	var val int64
	err := binary.Read(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (self *Decoder) Float() float32 {
	var val float32
	err := binary.Read(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (self *Decoder) Double() float64 {
	var val float64
	err := binary.Read(self, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
