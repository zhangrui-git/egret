package gate

import (
	"golang.org/x/net/websocket"
	"log"
	"net"
	"testing"
	"time"
)

func TestWebSocketClient(t *testing.T) {
	origin := "http://localhost:8080/ws"
	url := "ws://localhost:8080/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}

	//var msg = make([]byte, 512)
	//var n int
	//if n, err = ws.Read(msg); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Received: %s.\n", msg[:n])
	time.Sleep(30 * time.Second)
}

func TestTCPClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println(err.Error())
		return
	}
	n, err := conn.Write([]byte("hello world\n"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(n)
}
