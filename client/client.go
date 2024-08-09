package client

import (
	"sync"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/sentinel-official/sentinel-go-sdk/client/options"
	"github.com/sentinel-official/sentinel-go-sdk/types"
)

// Client contains necessary components for transaction handling, encoding, and decoding.
type Client struct {
	sync.Mutex                                // Mutex to ensure thread-safe access
	codec.ProtoCodecMarshaler                 // Marshaler for protobuf types
	client.TxConfig                           // Configuration for transactions
	kr                        keyring.Keyring // Keyring for managing keys
}

// New creates a new instance of Client with the provided ProtoCodecMarshaler.
// It returns a pointer to the newly created Client.
func New(protoCodec codec.ProtoCodecMarshaler) *Client {
	return &Client{
		ProtoCodecMarshaler: protoCodec,
		TxConfig:            authtx.NewTxConfig(protoCodec, authtx.DefaultSignModes),
	}
}

// NewDefault creates a new instance of Client with a default ProtoCodecMarshaler.
// It returns a pointer to the newly created Client.
func NewDefault() *Client {
	return New(codec.NewProtoCodec(types.NewInterfaceRegistry()))
}

// WithKeyring sets the Keyring for the Client and returns the updated Client instance.
func (c *Client) WithKeyring(kr keyring.Keyring) *Client {
	c.kr = kr
	return c
}

// Keyring returns the Keyring associated with the Client.
// If no Keyring was set using WithKeyring, it uses the provided Options to create one.
func (c *Client) Keyring(opts *options.Options) (keyring.Keyring, error) {
	if c.kr != nil {
		return c.kr, nil
	}

	// Create and return a new Keyring using options
	return opts.Keyring(c)
}
