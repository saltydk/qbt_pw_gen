package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go <password>")
		os.Exit(1)
	}

	password := os.Args[1]
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	iterations := 100000
	hash := pbkdf2.Key([]byte(password), salt, iterations, sha512.Size, sha512.New)

	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	hashBase64 := base64.StdEncoding.EncodeToString(hash)

	fmt.Printf("@ByteArray(%s:%s)\n", saltBase64, hashBase64)
}
