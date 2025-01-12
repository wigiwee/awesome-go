package main

import (
	"bytes"
	"fmt"
	"io"
)

type Conn struct {
	io.Writer
}

func NewConn() *Conn {
	return &Conn{
		Writer: new(bytes.Buffer),
	}
}

func (c *Conn) Write(b []byte) (int, error) {
	fmt.Println("Writing to underlying connection:", string(b))
	return c.Writer.Write(b)
}

type Server struct {
	peers []*Conn
}

func NewServer() *Server {
	s := &Server{
		peers: make([]*Conn, 0),
	}
	for i := 0; i < 10; i++ {
		s.peers = append(s.peers, NewConn())
	}
	return s
}

func (s *Server) broadcast(msg []byte) error {
	//using MultiWriter
	writer := make([]io.Writer, 0)
	for _, val := range s.peers {
		writer = append(writer, val)
	}
	mw := io.MultiWriter(writer...)
	_, err := mw.Write(msg)
	return err

	//Without usnig MultiWriter
	// for peer := range s.peers {
	// 	//use mutex here peers are mutually share resource
	// 	if _, err := peer.Write(msg); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// return nil
}

func main() {
	s := NewServer()
	s.broadcast([]byte("Foo"))
}
