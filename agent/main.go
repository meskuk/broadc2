package main

import (
	"github.com/meskuk/broadc2/internal"

	"fmt"
	"os"
	"encoding/base64"
	"crypto/ed25519"

	 _ "embed"
)

//go:embed fs/master.pub
var masterKeyB64 string
//go:embed fs/node.key
var nodeKeyB64 string

func main() {
	var masterKey ed25519.PublicKey
	masterKey, err := base64.StdEncoding.DecodeString(masterKeyB64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	nodeKeyBytes, err := base64.StdEncoding.DecodeString(nodeKeyB64)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	nodeKey := ed25519.NewKeyFromSeed(nodeKeyBytes)

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
