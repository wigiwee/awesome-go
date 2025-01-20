package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listneAddr string
	ln         net.Listener
	quitChan   chan struct{}
	msgch      chan Message
}

type Message struct {
	from    string
	payload []byte
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listneAddr: listenAddr,
		quitChan:   make(chan struct{}),
		msgch:      make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listneAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	s.ln = ln

	go s.acceptLoop()

	<-s.quitChan

	close(s.quitChan)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		fmt.Println("new connection received ", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	buf := make([]byte, 2048)
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}
		s.msgch <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}

		conn.Write([]byte("thank you\n"))
	}

}
func main() {
	server := NewServer(":3000")

	go func() {
		for msg := range server.msgch {
			fmt.Printf("received message %s : %s", msg.from, string(msg.payload))
		}
	}()
	log.Fatal(server.Start())

}
