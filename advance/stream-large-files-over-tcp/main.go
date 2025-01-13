package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type FileServer struct {
}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go fs.readLoop(conn)
	}
}

//non streaming way of sending file
// func (fs *FileServer) readLoop(conn net.Conn) {
// 	buf := make([]byte, 2048)
// 	for {
// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("received %d bytes over the network\n", n)
// 	}
// }

func (fs *FileServer) readLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	file := make([]byte, 0)
	for {
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			log.Fatal(err)
		}
		// panic("should panic")
		fmt.Println(buf.Bytes())
		file = append(file, buf.Bytes()...)
		fmt.Printf("received %d bytes over the network\n", n)
	}
}

func sendFile() error {
	//reading []byte / content of 20GB file
	fileContent, err := os.ReadFile("/home/happypotter/Videos/music/Taylor Swift Ready For It Audio_720pHF.mp4")
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}
	n, err := conn.Write(fileContent)
	if err != nil {
		return err
	}
	fmt.Printf("written %d bytes over the network\n", n)
	return nil
}

func streamFile() error {
	//reading []byte / content of 20GB file
	fileContent, err := os.ReadFile("/home/happypotter/Videos/music/Taylor Swift Ready For It Audio_720pHF.mp4")
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	binary.Write(conn, binary.LittleEndian, int64(len(fileContent)))
	n, err := io.Copy(conn, bytes.NewReader(fileContent))
	if err != nil {
		return err
	}
	fmt.Printf("written %d bytes over the network\n", n)
	return nil
}
func main() {
	//non streaming way
	//normal way to send file

	//sending file to the server
	// go func() {
	// 	time.Sleep(4 * time.Second)
	// 	sendFile()
	// }()

	// //starting server
	// server := &FileServer{}
	// server.start()

	//stream file

	go func() {
		time.Sleep(4 * time.Second)
		streamFile()
	}()

	//starting server
	server := &FileServer{}
	server.start()

}
