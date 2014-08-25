package PacketDecoder

import (
	"bytes"
	"encoding/binary"
)

func (self *Decoder) Packet() (*Decoder, error) {
	dataSize, err := self.Varint()
	if err != nil {
		return nil, err
	}

	data, err := self.ReadBytes(dataSize)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	return NewDecoder(reader), nil
}

func (self *Decoder) Varint() (uint64, error) {
	return binary.ReadUvarint(self)
}

func (self *Decoder) String() (string, error) {
	stringLength, err := self.Varint()
	if err != nil {
		return "", err
	}

	stringData, err := self.ReadBytes(stringLength)
	if err != nil {
		return "", err
	}

	return string(stringData), nil
}

func (self *Decoder) Bool() (bool, error) {
	b, err := self.Byte()
	if b != 0 && b != 1 && err == nil {
		err = InvalidBoolean
	}

	return (b == 1), err
}

func (self *Decoder) Byte() (byte, error) {
	return self.ReadByte()
}
func (self *Decoder) Short() (int16, error) {
	var val int16
	err := binary.Read(self.reader, binary.BigEndian, &val)
	return val, err
}
func (self *Decoder) Ushort() (uint16, error) {
	var val uint16
	err := binary.Read(self.reader, binary.BigEndian, &val)
	return val, err
}
func (self *Decoder) Int() (int32, error) {
	var val int32
	err := binary.Read(self.reader, binary.BigEndian, &val)
	return val, err
}
func (self *Decoder) Long() (int64, error) {
	var val int64
	err := binary.Read(self.reader, binary.BigEndian, &val)
	return val, err
}
func (self *Decoder) Float() (float32, error) {
	var val float32
	err := binary.Read(self.reader, binary.BigEndian, &val)
	return val, err
}
func (self *Decoder) Double() (float64, error) {
	var val float64
	err := binary.Read(self.reader, binary.BigEndian, &val)
	return val, err
}
