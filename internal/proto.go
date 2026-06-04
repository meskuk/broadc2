package internal

import (
	"crypto/ed25519"
	"encoding/base64"
	"slices"
	"fmt"
)

// A message.
type Message struct {
	Version int // 1 byte version code. Currently at 1.
	Length int // Length
	Signature []byte // 64 byte ed25519 Signature
	Content []byte // Content
}

type VerificationError error

// Unmarshal from raw bytes into struct fields
func (m *Message) Unmarshal(data []byte, key ed25519.PublicKey) (error) {
	// TODO. FIXME: !!!: THISSUCKS: This looks complicated and will violently explode
	// when I try to add more fields. I think using buffers and the struct fields directly
	// would work well. 
	m.Version = int(data[0])
	m.Length = int(data[1])
	contentStart := ed25519.SignatureSize+2
	m.Signature = data[2:contentStart]
	m.Content = data[contentStart:contentStart+m.Length]
	if !ed25519.Verify(key, m.Content, m.Signature) {
		return fmt.Errorf(
			"Packet failed verification: signature: %s",
			base64.StdEncoding.EncodeToString(m.Signature),
		)
	}
	return nil
}

// Marshal from struct into raw bytes
func (m *Message) Marshal() []byte {
	bytes := slices.Concat([]byte{1, byte(m.Length)}, m.Signature, m.Content)
	return bytes
}
