package Packets

import (
	"encoding/binary"
)

func (s *Stream) writeUvarint(v uint64) {
	s.writeBytes(makeUvarint(v))
}

func (s *Stream) writeString(str string) {
	stringLength := uint64(len(str))
	s.writeUvarint(stringLength)
	s.writeBytes([]byte(str))
}

func (s *Stream) bool(val bool) {
	var b byte
	if val {
		b = 1
	} else {
		b = 0
	}
	s.writeByte(b)
}
func (s *Stream) writeBytes(p []byte) {
	n, err := s.outBuffer.Write(p)
	if err != nil {
		panic(err)
	}
	if n != len(p) {
		panic("could not write to buffer")
	}
}
func (s *Stream) writeByte(val byte) {
	s.writeBytes([]byte{val})
}

func (s *Stream) writeShort(val int16) {
	err := binary.Write(s.outBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (s *Stream) writeUshort(val uint16) {
	err := binary.Write(s.outBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (s *Stream) writeInt(val int32) {
	err := binary.Write(s.outBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (s *Stream) writeLong(val int64) {
	err := binary.Write(s.outBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (s *Stream) writeFloat(val float32) {
	err := binary.Write(s.outBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
func (s *Stream) writeDouble(val float64) {
	err := binary.Write(s.outBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
}
