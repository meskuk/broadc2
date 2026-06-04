package main

import (
	"github.com/meskuk/broadc2/internal"

	"fmt"
	"os"
	"encoding/base64"
	"crypto/ed25519"
)

func main() {
	var masterKey ed25519.PublicKey
	mKeyRaw, err := os.ReadFile("master.pub")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	masterKey, err = base64.StdEncoding.DecodeString(string(mKeyRaw))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	nKeyRaw, err := os.ReadFile("node.key")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	nKeyBytes, err := base64.StdEncoding.DecodeString(string(nKeyRaw))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	nodeKey := ed25519.NewKeyFromSeed(nKeyBytes)

	s := internal.Server{
		NodeKey: nodeKey,
		MasterKey: masterKey,
	}
	ch := make(chan internal.Message)
	go s.Listen(ch)
	fmt.Println("Started listener")

	// Print received messages
	for {
		msg := <-ch
		fmt.Println("Message:", string(msg.Content))
	}
}
