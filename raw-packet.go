package main

import (
	"./packet-decoder"
	"bytes"
	"encoding/binary"
	"io"
)

type RawPacket struct {
	Payload []byte
}

func NextPacket(src io.Reader) (*RawPacket, error) {
	result := new(RawPacket)
	decoder := PacketDecoder.NewDecoder(src)

	length, err := decoder.Varint()
	if err != nil {
		return nil, err
	}

	result.Payload, err = decoder.ReadBytes(length)
	if err != nil {
		return nil, err
	}

	return result, err
}

/*


type ByteReader interface {
	Read([]byte) (int, error)
	ReadByte() (byte, error)
}

func ReadVarint(reader io.ByteReader) (int32, error) {
	shift := byte(0)
	result := int32(0)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			return result, err
		}
		result |= int32(b&0x7f) << shift

		if (b & 0x80) == 0 {
			return result, nil
		}
		shift += 7

		if shift >= 64 {
			panic("varint is too big")
		}
	}
}

var varintParseError = errors.New("varint could not be parsed")
var streamClosed = errors.New("stream closed")

func NextPacket(src ByteReader) (*RawPacket, error) {
	result := new(RawPacket)

	decoder := PacketDecoder

	packetSize, err := binary.ReadUvarint(src)
	if err != nil {
		return result, err
	}

	data := make([]byte, packetSize)
	read := 0
	for read < len(data) {
		n, err := src.Read(data)
		if n == 0 {
			return result, streamClosed
		}
		if err != nil {
			return result, err
		}
		read += n
	}

	packetType, n := binary.Uvarint(data)
	if n <= 0 {
		return result, varintParseError
	}
	result.Type = packetType

	result.Payload = data[n:]

	return result, nil
}
*/

func (p *RawPacket) Serialize() []byte {
	lenData := make([]byte, binary.MaxVarintLen32)
	totalLength := uint64(len(p.Payload))
	lenLen := binary.PutUvarint(lenData, totalLength)

	data := [][]byte{lenData[:lenLen], p.Payload}

	return bytes.Join(data, []byte{})
}
