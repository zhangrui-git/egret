package gate

import (
	"github.com/zhangrui-git/egret/message"
	"testing"
)

func TestGateway_StartTCPServer(t *testing.T) {
	type fields struct {
		addr string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "test1", fields: fields{addr: "localhost:8080"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gw := NewGateway(tt.fields.addr)
			gw.StartTCPServer()
		})
	}
}

func TestGateway_StartWebSocketServer(t *testing.T) {
	type fields struct {
		addr string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "test1", fields: fields{addr: ":8080"}, args: args{path: "/ws"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gw := NewGateway(tt.fields.addr)
			gw.Codec(message.JsonEncoder, message.JsonDecoder)
			gw.StartWebSocketServer(tt.args.path)
		})
	}
}
