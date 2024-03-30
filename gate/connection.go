package gate

import (
	"bufio"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"sync"
)

type Connection interface {
	Serve()
	Close() error
}

type WSConnection struct {
	gateway *Gateway
	conn    *websocket.Conn
}

func NewWSConnection(gw *Gateway, conn *websocket.Conn) *WSConnection {
	return &WSConnection{gateway: gw, conn: conn}
}

func (wc *WSConnection) Serve() {
	for {
		rAddr := wc.conn.RemoteAddr()
		log.Println(rAddr.String())

		_, msg, err := wc.conn.ReadMessage()
		if err != nil {
			switch eType := err.(type) {
			default:
				log.Println(eType)
				log.Println(err.Error())
				break
			}
			/*if websocket.IsCloseError(err) {
				log.Println("连接断开")
				break
			}*/
			break
		}
		log.Println("新的消息")
		log.Println(wc.gateway.Decode(msg))
	}
}

func (wc *WSConnection) Close() error {
	return wc.conn.Close()
}

type TCPConnection struct {
	gateway *Gateway
	conn    net.Conn
}

func NewTCPConnection(gw *Gateway, conn net.Conn) *TCPConnection {
	return &TCPConnection{gateway: gw, conn: conn}
}

func (tc *TCPConnection) Serve() {
	reader := bufio.NewReader(tc.conn)
	for {
		rAddr := tc.conn.RemoteAddr()
		log.Println(rAddr.String())

		msg := make([]byte, 1024)
		n, err := reader.Read(msg)
		if err != nil {
			/*if errors.Is(err, net.ErrClosed) {
				log.Println("连接断开")
				break
			}*/
			var e net.Error
			switch {
			case errors.As(err, &e):
				log.Println("net err", err.Error())
				break
			default:
				log.Println(err.Error())
				break
			}
			break
		}
		log.Println("新的消息")
		log.Println(tc.gateway.Decode(msg[:n]))
	}
}

func (tc *TCPConnection) Close() error {
	return tc.conn.Close()
}

type ConnectionSet struct {
	mu   sync.RWMutex
	conn map[string]Connection
}

func NewConnectionSet() *ConnectionSet {
	return &ConnectionSet{conn: make(map[string]Connection)}
}

func (cSet *ConnectionSet) AddConn(pk string, conn Connection) {
	cSet.mu.Lock()
	defer cSet.mu.Unlock()
	cSet.conn[pk] = conn
	log.Println("add conn:", pk, len(cSet.conn))
}

func (cSet *ConnectionSet) RemoveConn(pk string) {
	cSet.mu.Lock()
	defer cSet.mu.Unlock()
	conn, ok := cSet.conn[pk]
	if ok {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
		delete(cSet.conn, pk)
	}
}

func (cSet *ConnectionSet) IterateConn(f func(conn Connection) bool) {
	cSet.mu.RLock()
	defer cSet.mu.RUnlock()
	for pk, conn := range cSet.conn {
		log.Println("iterate conn:", pk)
		if !f(conn) {
			break
		}
	}
}
