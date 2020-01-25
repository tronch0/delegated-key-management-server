package main

import (
	"../key-management-delegated-keys/client"
	"../key-management-delegated-keys/crypto"
	"../key-management-delegated-keys/server"

	"fmt"
)

func main() {

	fmt.Println("Creating Client and two Servers...")
	client := createClient()
	serverA := createServer()
	serverB := createServer()

	password := "“Privacy is not something that I’m merely entitled to, it’s an absolute prerequisite.” -M.B"
	passwordB := []byte(password)
	fmt.Printf("Client password: %s\n", password)

	fmt.Println("Client hide his commitment to password (Client hash-to-curve then exponentiate to one time encryption key")
	xWithClientSecret, yWithClientSecret := client.Hide(passwordB)

	fmt.Println("Servers A and B applying their keys to the output of the client")
	xWithClientAndFirstServerSecrets, yWithClientAndFirstServerSecrets := serverA.ApplyKey(xWithClientSecret, yWithClientSecret)
	xWithClientAndBothServersSecrets, yWithClientAndBothServersSecrets := serverB.ApplyKey(xWithClientAndFirstServerSecrets, yWithClientAndFirstServerSecrets)

	fmt.Println("Client unhide his commitment (Client exponentiate to the inverse of his initial one time encryption key")
	xWithServersSecrets, yWithServersSecrets := client.Unhide(xWithClientAndBothServersSecrets, yWithClientAndBothServersSecrets)

	fmt.Printf("\n Result:\n\tx-%x \n\ty-%x \n", xWithServersSecrets, yWithServersSecrets)
}

func createServer() *server.Server {
	serverSk, err := crypto.GenerateR()
	if err != nil {
		panic("error on crypto.GenerateR()")
	}
	return server.New(serverSk)
}

func createClient() *client.Client {
	c, err := client.New()
	if err != nil {
		panic("error on client ctor")
	}

	return c
}
