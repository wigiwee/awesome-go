package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {

	payload := []byte("Hello from go programming language")
	hashAndBroadcast(newHashReader(payload))
}

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	bytes.Reader //this means that hashReader will inherit all the methods of bytes.Readers
	buf          *bytes.Buffer
}

func newHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: *bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r HashReader) error {
	// hash := r.(*hashReader).hash()
	hash := r.hash()
	fmt.Println(hash)
	return broadcast(r)

}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("String of the bytes is ", string(b))
	return nil
}
