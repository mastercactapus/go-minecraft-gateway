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
func (self *Decoder) binaryRead(v interface{}) interface{} {
	err := binary.Read(self, binary.BigEndian, &v)
	if err != nil {
		panic(err)
	}
	return v
}
func (self *Decoder) Short() int16 {
	var val int16
	return self.binaryRead(val).(int16)
}
func (self *Decoder) Ushort() uint16 {
	var val uint16
	return self.binaryRead(val).(uint16)
}
func (self *Decoder) Int() int32 {
	var val int32
	return self.binaryRead(val).(int32)
}
func (self *Decoder) Long() int64 {
	var val int64
	return self.binaryRead(val).(int64)
}
func (self *Decoder) Float() float32 {
	var val float32
	return self.binaryRead(val).(float32)
}
func (self *Decoder) Double() float64 {
	var val float64
	return self.binaryRead(val).(float64)
}
