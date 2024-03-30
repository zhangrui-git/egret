package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

type user struct {
	Name string
	Age  uint8
}

var p user

func main() {
	fmt.Printf("%+v\n", p)
	http.Handle("/", websocket.Handler(func(conn *websocket.Conn) {
		log.Println("新的连接")
	}))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	/*
		tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8081")
		if err != nil {
			return
		}
		tcpListener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			return
		}
		defer func(tcpListener *net.TCPListener) {
			err := tcpListener.Close()
			if err != nil {
				return
			}
		}(tcpListener)
		for {
			tcpConn, err := tcpListener.AcceptTCP()
			if err != nil {
				return
			}
			go func(conn *net.TCPConn) {
				b := make([]byte, 1024)
				buf := bytes.NewBuffer(b)
				i, err := io.Copy(buf, conn)
				if err != nil {
					return
				}
				if i > 0 {
					fmt.Println(b)
				}
			}(tcpConn)
		}*/
}
