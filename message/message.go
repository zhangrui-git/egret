package message

import "encoding/json"

type UplinkMsg struct {
	Header struct {
		MessageId      string `json:"msgid"`
		ServiceCommand string `json:"srvcmd"`
		Version        int    `json:"ver"`
		RequestTime    int    `json:"reqtime"`
		Platform       int    `json:"platform"`
	} `json:"header"`
	Body json.RawMessage `json:"body"`
}

type DownlinkMsg struct {
	Header struct {
		MessageId      string `json:"msgid"`
		ServiceCommand string `json:"srvcmd"`
		ErrorCode      int    `json:"errcode"`
		ErrorMessage   string `json:"errmsg"`
	} `json:"header"`
	Body json.RawMessage `json:"body"`
}
