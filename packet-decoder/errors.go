package PacketDecoder

import "errors"

var EndOfStream = errors.New("end of stream")
var InvalidBoolean = errors.New("invalid value for boolean")
