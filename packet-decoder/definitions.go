package PacketDecoder

type MultiReader interface {
	Read([]byte) (int, error)
	ReadByte() (byte, error)
}

type Decoder struct {
	reader MultiReader
}

func NewDecoder(reader MultiReader) *Decoder {
	return &Decoder{reader}
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
