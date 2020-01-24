package main

import (
	"../key-management-delegated-keys/client"
	"../key-management-delegated-keys/crypto"
	"../key-management-delegated-keys/server"
	"fmt"
	"log"
)

func main() {

	data := "“Privacy is not something that I’m merely entitled to, it’s an absolute prerequisite.” -M.B"
	dataB := []byte(data)

	c, err := client.New()
	if err != nil {
		log.Fatal("error on client ctor")
	}

	clientX, clientY := c.Hide(dataB)

	serverSk, err := crypto.GenerateR()
	if err != nil {
		log.Fatal("error on crypto.GenerateR()")
	}

	s := server.New(serverSk)

	xWithServerKey, yWithServerKey := s.ApplyKey(clientX, clientY)

	x, y := c.Unhide(xWithServerKey, yWithServerKey)

	fmt.Printf("\n x: %x \n y: %x \n", x, y)
}
