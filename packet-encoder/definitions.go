package PacketEncoder

import (
	"bytes"
)

const (
	MaxPacketSize = 1024 * 1024 * 32
)

type MultiWriter interface {
	Write([]byte) (int, error)
	WriteByte(byte) error
}

type Encoder struct {
	Writer  MultiWriter
	Written uint64
	Data    []byte
}

func NewEncoder(writer MultiWriter) *Encoder {
	return &Encoder{writer, 0, []byte{}}
}

func NewPacket() *Encoder {
	enc := new(Encoder)
	enc.Data = make([]byte, MaxPacketSize)
	writer := MultiWriter(bytes.NewBuffer(enc.Data))
	enc.Writer = writer
	enc.Written = 0
	return enc
}

func (self *Encoder) Write(data []byte) (int, error) {
	n, err := self.Writer.Write(data)
	self.Written += uint64(n)
	return n, err
}
func (self *Encoder) WriteByte(data byte) error {
	err := self.Writer.WriteByte(data)
	self.Written++
	return err
}

func (self *Encoder) WriteBytes(data []byte) error {
	wroteTotal := 0
	total := len(data)
	for wroteTotal < total {
		wrote, err := self.Write(data[wroteTotal:])
		if err != nil {
			return err
		}

		wroteTotal += wrote
	}

	return nil
}
