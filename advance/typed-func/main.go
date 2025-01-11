package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunc func(string) string

type Server struct {
	filenameTransformerFunc TransformFunc
}

func main() {
	s := &Server{
		filenameTransformerFunc: addCustomFilenamePrefix("ABC_"),
	}
	s.handleRequest("cool.go")
}

func KKprefixToFilename(filename string) string {
	// hash := sha256.Sum256([]byte(filename))
	// newFilename := hex.EncodeToString(hash[:])
	return "KK_" + filename
}

func addCustomFilenamePrefix(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}
func hashFilenameSha1(filename string) string {
	hash := sha1.New().Sum([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}

func hashFilenameSha256(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}

// maybe you want of calculate hash differently for diff server
// sha1
// prefix KK_
// hmac

func (s *Server) handleRequest(filename string) error {
	newFilename := s.filenameTransformerFunc(filename)
	fmt.Println("new file name : ", newFilename)
	return nil
}
