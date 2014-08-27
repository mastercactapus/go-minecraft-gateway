package Packets

import (
	"encoding/binary"
	"io"
)

func (s *Stream) readUvarint() uint64 {
	n, err := binary.ReadUvarint(s.inBuffer)
	if err != nil {
		panic(err)
	}
	return n
}

func (s *Stream) readString() string {
	stringLength := s.readUvarint()
	stringData := s.readBytes(stringLength)
	return string(stringData)
}

func (s *Stream) readBool() bool {
	b := s.readByte()
	if b != 0 && b != 1 {
		panic("invalid boolean value")
	}

	return (b == 1)
}

func (s *Stream) readBytes(n uint64) []byte {
	buf := make([]byte, n)
	numRead, err := io.ReadFull(s.inBuffer, buf)
	if err != nil {
		panic(err)
	}
	if uint64(numRead) != n {
		panic("could not read from buffer")
	}
	return buf
}
func (s *Stream) readByte() byte {
	b, err := s.inBuffer.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

func (s *Stream) readShort() int16 {
	var val int16
	err := binary.Read(s.inBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (s *Stream) readUshort() uint16 {
	var val uint16
	err := binary.Read(s.inBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (s *Stream) readInt() int32 {
	var val int32
	err := binary.Read(s.inBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (s *Stream) readLong() int64 {
	var val int64
	err := binary.Read(s.inBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (s *Stream) readFloat() float32 {
	var val float32
	err := binary.Read(s.inBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
func (s *Stream) readDouble() float64 {
	var val float64
	err := binary.Read(s.inBuffer, binary.BigEndian, &val)
	if err != nil {
		panic(err)
	}
	return val
}
