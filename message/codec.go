package message

import "encoding/json"

type Encoder interface {
	Encode(*DownlinkMsg) ([]byte, error)
}

type Decoder interface {
	Decode([]byte) (*UplinkMsg, error)
}

type EncodeFunc func(*DownlinkMsg) ([]byte, error)

func (f EncodeFunc) Encode(msg *DownlinkMsg) ([]byte, error) {
	return f(msg)
}

type DecodeFunc func([]byte) (*UplinkMsg, error)

func (f DecodeFunc) Decode(data []byte) (*UplinkMsg, error) {
	return f(data)
}

var JsonEncoder EncodeFunc = func(msg *DownlinkMsg) ([]byte, error) {
	return json.Marshal(msg)
}

var JsonDecoder DecodeFunc = func(data []byte) (*UplinkMsg, error) {
	msg := &UplinkMsg{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
