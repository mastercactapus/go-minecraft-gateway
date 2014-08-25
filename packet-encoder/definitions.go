package PacketEncoder

import (
	"bytes"
	"io"
)

type Encoder struct {
	writer      io.Writer
	written     uint64
	bytesBuffer *bytes.Buffer
}

func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{writer, 0, new(bytes.Buffer)}
}

func NewPacket() *Encoder {
	enc := new(Encoder)
	enc.bytesBuffer = new(bytes.Buffer)
	enc.writer = enc.bytesBuffer
	enc.written = 0
	return enc
}

func (self *Encoder) Write(data []byte) (int, error) {
	n, err := self.writer.Write(data)
	self.written += uint64(n)
	return n, err
}

func (self *Encoder) WriteByte(v byte) error {
	n, err := self.Write([]byte{v})
	if n == 0 {
		return io.EOF
	}
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

func (self *Encoder) writeBytePanic(v byte) {
	err := self.WriteByte(v)
	if err != nil {
		panic(err)
	}
}

func (self *Encoder) writeBytesPanic(v []byte) {
	err := self.WriteBytes(v)
	if err != nil {
		panic(err)
	}
}
