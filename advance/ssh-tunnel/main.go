package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
)

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectSSH(username, password, host string, port int) (*ssh.Session, *ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial: %w", err)
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create session: %w", err)
	}

	return session, client, nil
}
func streamCommand(session *ssh.Session, command string, outputChan chan string) error {
	stdout, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	err = session.Start(command)
	if err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if n > 0 {
			outputChan <- string(buf[:n]) // Stream output
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading output: %w", err)
		}
	}

	close(outputChan)
	return session.Wait()
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade failed:", err)
		return
	}
	defer conn.Close()

	// Placeholder for SSH credentials
	_, usernameBytes, err := conn.ReadMessage()
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid Username"))
		log.Println("error parsing username")
		return
	}
	username := string(usernameBytes)
	_, passwordBytes, err := conn.ReadMessage()
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid password"))
		log.Println("error parsing password")
		return
	}
	password := string(passwordBytes)

	_, hostBytes, err := conn.ReadMessage()
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("invalid host"))
		log.Println("Invalid host addr")
		return
	}
	host := string(hostBytes)

	port := 22
	for {
		session, client, err := connectSSH(username, password, host, port)
		if err != nil {
			log.Println("SSH Connection failed:", err)
			return
		}
		outputChan := make(chan string)

		_, command, err := conn.ReadMessage()
		log.Println("command received", command)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("error reading Command"))
			continue
		}
		go func() {
			err := streamCommand(session, string(command), outputChan)
			if err != nil {
				log.Println("Command execution failed:", err)
			}
		}()

		// Send output to WebSocket
		for output := range outputChan {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(output)); err != nil {
				log.Println("WebSocket write failed:", err)
				break
			}
		}
		client.Close()
		session.Close()
	}
}
