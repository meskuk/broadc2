package main

import (
	"github.com/meskuk/broadc2/internal"

	"fmt"
	"flag"
	"net"
	"os"
	
	"encoding/base64"
	"crypto/ed25519"
)

func send(content []byte, key ed25519.PrivateKey) {
	// This function should send the string to all servers by making an [internal.Message],
	// marshalling it, then sending.
	sig := ed25519.Sign(key, content)
	msg := internal.Message{
		Version: 1,
		Length: len(content),
		Signature: sig,
		Content: content,
	}
	fmt.Println(base64.StdEncoding.EncodeToString(sig))
	bytes := msg.Marshal()
	c, err := net.Dial("udp4", "255.255.255.255:22222")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	c.Write(bytes)
	fmt.Println("Sent", len(bytes))
}

func main() {
	// Flags:
	//   -send <msg> - broadcast a message to nodes
	//   -generate - generate a 'master key' and a 'node key' to embed in binaries
	sendContent := flag.String("send", "", "Message to send")
	generateFlag := flag.Bool("generate", false, "Generate keys")
	flag.Parse()
	if *sendContent != "" {
		// Read in the seed
		seedBytes, err := os.ReadFile("master.key")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		// Decode it and get the key
		seed, err := base64.StdEncoding.DecodeString(string(seedBytes))
		masterKey := ed25519.NewKeyFromSeed(seed)
		// Send t
		send([]byte(*sendContent), masterKey)
	} else if *generateFlag {
		fmt.Println("Generating master key")
		pub, priv, _ := ed25519.GenerateKey(nil)
		encodedPriv := base64.StdEncoding.EncodeToString(priv.Seed())
		encodedPub := base64.StdEncoding.EncodeToString(pub)
		err := os.WriteFile("master.key", []byte(encodedPriv), 0600)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.WriteFile("master.pub", []byte(encodedPub), 0644)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Generating node key")
		nodePub, nodePriv, _ := ed25519.GenerateKey(nil)
		encodedNodePriv := base64.StdEncoding.EncodeToString(nodePriv.Seed())
		encodedNodePub := base64.StdEncoding.EncodeToString(nodePub)
		err = os.WriteFile("node.key", []byte(encodedNodePriv), 0600)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.WriteFile("node.pub", []byte(encodedNodePub), 0644)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
