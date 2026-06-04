package internal

import (
	"fmt"
	"net"
	"crypto/ed25519"
	"golang.org/x/net/ipv4"
)

type Server struct {
	NodeKey ed25519.PrivateKey
	MasterKey ed25519.PublicKey
	packetconn *ipv4.PacketConn
}

func (s *Server) Send(msg Message) {
	data := msg.Marshal()
	s.packetconn.WriteTo(
		data,
		nil,
		&net.UDPAddr{IP: net.IPv4(255,255,255,255), Port: 22222},
	)
}

func (s *Server) Listen(ch chan Message) (error) {
	c, err := net.ListenPacket("udp4", "0.0.0.0:22222")
	if err != nil {
		return err
	}
	defer c.Close()

	packetconn := ipv4.NewPacketConn(c)
	s.packetconn = packetconn

	// TODO: I should read in a way that grows the slice
	buf := make([]byte, 1024)
	for {
		_, _, src, err := packetconn.ReadFrom(buf)
		if err != nil {
			// TODO: Um. When I stop debugging with println, do I just ignore errors here?
			fmt.Println("Error accepting packet:", err.Error())
		}
		fmt.Println("Reading packet from", src)
		var msg Message
		err = msg.Unmarshal(buf, s.MasterKey)
		if err != nil {
			fmt.Println("Error unmarshalling packet:", err.Error())
		}
		ch<-msg
		
	}
}
