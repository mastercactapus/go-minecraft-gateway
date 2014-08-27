package Packets

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

type Stream struct {
	stream     *bufio.ReadWriter
	inBuffer   *bytes.Buffer
	outBuffer  *bytes.Buffer
	hasPacket  bool
	packetType uint64
}

func NewStream(stream io.ReadWriter) *Stream {
	s := new(Stream)
	s.inBuffer = new(bytes.Buffer)
	s.outBuffer = new(bytes.Buffer)
	reader := bufio.NewReader(stream)
	writer := bufio.NewWriter(stream)
	s.stream = bufio.NewReadWriter(reader, writer)
	s.hasPacket = false
	return s
}

func makeUvarint(v uint64) []byte {
	buf := make([]byte, binary.MaxVarintLen32)
	size := binary.PutUvarint(buf, v)
	return buf[:size]
}

func (s *Stream) PeekPacketType() uint64 {
	if s.hasPacket {
		return s.packetType
	}

	s.inBuffer.Reset()
	size, err := binary.ReadUvarint(s.stream)
	if err != nil {
		panic(err)
	}
	n, err := io.CopyN(s.inBuffer, s.stream, int64(size))
	if err != nil {
		panic(err)
	}
	if uint64(n) != size {
		panic("could not read from stream")
	}

	s.packetType = s.readUvarint()
	s.hasPacket = true
	return s.packetType
}

func (s *Stream) writePacket() {
	size := s.outBuffer.Len()
	varint := makeUvarint(uint64(size))
	n, err := s.stream.Write(varint)
	if err != nil {
		panic(err)
	}
	if n != len(varint) {
		panic("could not write to stream")
	}

	numSent, err := io.CopyN(s.stream, s.outBuffer, int64(size))
	if err != nil {
		panic(err)
	}
	if numSent != int64(size) {
		panic("could not write to stream")
	}

	s.outBuffer.Reset()
}
