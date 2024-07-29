package wireguard

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/curve25519"
)

const KeyLength = 32

// Key represents a 32-byte key used in WireGuard.
type Key [KeyLength]byte

// String returns the base64 encoding of the key.
func (k *Key) String() string {
	return base64.StdEncoding.EncodeToString(k[:])
}

// IsZero checks if the key is all zeros.
func (k *Key) IsZero() bool {
	var zeros Key
	return subtle.ConstantTimeCompare(k[:], zeros[:]) == 1
}

// Public calculates the public key from a private key.
func (k *Key) Public() *Key {
	var pub [KeyLength]byte
	curve25519.ScalarBaseMult(&pub, (*[KeyLength]byte)(k))
	return (*Key)(&pub)
}

// NewPresharedKey generates a new random 32-byte key.
func NewPresharedKey() (*Key, error) {
	var k Key
	if _, err := rand.Read(k[:]); err != nil {
		return nil, fmt.Errorf("failed to generate preshared key: %w", err)
	}
	return &k, nil
}

// NewPrivateKey generates a new private key with the required properties.
func NewPrivateKey() (*Key, error) {
	k, err := NewPresharedKey()
	if err != nil {
		return nil, err
	}
	k[0] &= 248
	k[31] = (k[31] & 127) | 64
	return k, nil
}

// NewPrivateKeyFromString decodes a base64-encoded string to a Key.
func NewPrivateKeyFromString(s string) (*Key, error) {
	v, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("invalid base64 encoding: %w", err)
	}
	if len(v) != KeyLength {
		return nil, errors.New("decoded key must be 32 bytes")
	}

	var key Key
	copy(key[:], v)
	return &key, nil
}
