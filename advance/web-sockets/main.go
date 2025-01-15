package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func newServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client", ws.RemoteAddr())

	//use mutex for conns to avoid race conditions
	s.conns[ws] = true

	s.readLoop(ws)
	return
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Read error :", err)
			continue
		}
		msg := buf[:n]
		// fmt.Println(string(msg))
		// ws.Write([]byte("Thanks for the message"))

		s.braodcast(msg)
	}

}

func (s *Server) handleWsOrderbook(ws *websocket.Conn) {

	fmt.Println("new incoming connection to client to orderbook feed", ws.RemoteAddr())
	for {
		paylod := fmt.Sprintf("orderbook data ->%d \n", time.Now().UnixNano())
		ws.Write([]byte(paylod))
		time.Sleep(time.Second * 2)
	}

}
func (s *Server) braodcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error ", err)
			}
		}(ws)
	}
}
func main() {
	s := newServer()
	http.Handle("/ws", websocket.Handler(s.handleWS))
	http.Handle("/orderbookFeed", websocket.Handler(s.handleWsOrderbook))
	http.ListenAndServe(":3000", nil)

}
