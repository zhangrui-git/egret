package packet

type Encoder interface {
	Encode(v any) ([]byte, error)
}

type Decoder interface {
	Decode(v []byte) (any, error)
}
