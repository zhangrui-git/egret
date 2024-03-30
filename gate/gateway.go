package gate

import (
	"github.com/gorilla/websocket"
	"github.com/zhangrui-git/egret/message"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Gateway struct {
	addr    string
	connSet *ConnectionSet
	encoder message.Encoder
	decoder message.Decoder
}

func NewGateway(addr string) *Gateway {
	return &Gateway{
		addr:    addr,
		connSet: NewConnectionSet(),
		encoder: message.JsonEncoder,
		decoder: message.JsonDecoder,
	}
}

func (gw *Gateway) Codec(encoder message.Encoder, decoder message.Decoder) {
	gw.encoder = encoder
	gw.decoder = decoder
}

func (gw *Gateway) Encode(msg *message.DownlinkMsg) ([]byte, error) {
	return gw.encoder.Encode(msg)
}

func (gw *Gateway) Decode(data []byte) (*message.UplinkMsg, error) {
	return gw.decoder.Decode(data)
}

func (gw *Gateway) StartWebSocketServer(path string) {
	log.Println("start websocket gateway server")

	upgrade := websocket.Upgrader{}
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		log.Println("新的连接")
		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err.Error())
			return
		}
		wsConn := NewWSConnection(gw, conn)
		ts := time.Now().UnixMicro()
		gw.connSet.AddConn(strconv.FormatInt(ts, 10), wsConn)
		go wsConn.Serve()
	})
	if err := http.ListenAndServe(gw.addr, nil); err != nil {
		log.Println(err.Error())
	}
}

func (gw *Gateway) StartTCPServer() {
	log.Println("start TCP gateway server")

	listener, err := net.Listen("tcp", gw.addr)
	if err != nil {
		log.Println(err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		log.Println("新的连接")

		tcpConn := NewTCPConnection(gw, conn)
		ts := time.Now().UnixMicro()
		gw.connSet.AddConn(strconv.FormatInt(ts, 10), tcpConn)
		go tcpConn.Serve()
	}
}
