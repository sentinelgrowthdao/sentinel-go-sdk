package conf

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"net/netip"
	"strings"

	"golang.org/x/crypto/curve25519"
)

const KeyLength = 32

type Endpoint struct {
	Host string
	Port uint16
}

func (e *Endpoint) String() string {
	if strings.IndexByte(e.Host, ':') != -1 {
		return fmt.Sprintf("[%s]:%d", e.Host, e.Port)
	}
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

func (e *Endpoint) IsEmpty() bool {
	return len(e.Host) == 0
}

type Key [KeyLength]byte

func (k *Key) String() string {
	return base64.StdEncoding.EncodeToString(k[:])
}

func (k *Key) IsZero() bool {
	var zeros Key
	return subtle.ConstantTimeCompare(zeros[:], k[:]) == 1
}

func (k *Key) Public() *Key {
	var p [KeyLength]byte
	curve25519.ScalarBaseMult(&p, (*[KeyLength]byte)(k))
	return (*Key)(&p)
}

func NewPresharedKey() (*Key, error) {
	var k [KeyLength]byte
	_, err := rand.Read(k[:])
	if err != nil {
		return nil, err
	}
	return (*Key)(&k), nil
}

func NewPrivateKey() (*Key, error) {
	k, err := NewPresharedKey()
	if err != nil {
		return nil, err
	}
	k[0] &= 248
	k[31] = (k[31] & 127) | 64
	return k, nil
}

func NewPrivateKeyFromString(b64 string) (*Key, error) {
	return parseKeyBase64(b64)
}

type Interface struct {
	PrivateKey Key
	Addresses  []netip.Prefix
	ListenPort uint16
	MTU        uint16
	DNS        []netip.Addr
	DNSSearch  []string
	PreUp      string
	PostUp     string
	PreDown    string
	PostDown   string
	TableOff   bool
}

type Peer struct {
	PublicKey           Key
	PresharedKey        Key
	AllowedIPs          []netip.Prefix
	Endpoint            Endpoint
	PersistentKeepalive uint16
}

type Config struct {
	Name      string
	Interface Interface
	Peers     []Peer
}
