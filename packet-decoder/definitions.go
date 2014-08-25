package PacketDecoder

import (
	"io"
)

type Decoder struct {
	reader io.Reader
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{reader}
}

func (self *Decoder) Read(p []byte) (int, error) {
	return self.reader.Read(p)
}

func (self *Decoder) ReadByte() (byte, error) {
	b := make([]byte, 1)
	n, err := self.reader.Read(b)
	if err != nil {
		return 0, err
	}
	if n == 0 {
		return 0, io.EOF
	}

	return b[0], nil
}

func (self *Decoder) ReadBytes(n uint64) ([]byte, error) {
	read := uint64(0)
	buf := make([]byte, n)
	for read < n {
		count, err := self.reader.Read(buf[read:])
		if count == 0 {
			return buf, EndOfStream
		}
		if err != nil {
			return buf, err
		}
		read += uint64(count)
	}

	return buf, nil
}

func (self *Decoder) readBytesPanic(n uint64) []byte {
	data, err := self.ReadBytes(n)
	if err != nil {
		panic(err)
	}
	return data
}

func (self *Decoder) readBytePanic() byte {
	b, err := self.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}
