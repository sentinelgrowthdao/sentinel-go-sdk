package client

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/sentinel-official/sentinel-go-sdk/v1/client/options"
)

// Key retrieves key information from the keyring based on the provided name and options.
func (c *Context) Key(name string, opts *options.KeyOptions) (keyring.Info, error) {
	// Initialize a keyring based on the provided options.
	kr, err := opts.Keyring()
	if err != nil {
		return nil, err
	}

	// Retrieve key information from the keyring.
	return kr.Key(name)
}

// Sign signs the provided data using the key from the keyring specified by the name and options.
func (c *Context) Sign(name string, buf []byte, opts *options.KeyOptions) ([]byte, cryptotypes.PubKey, error) {
	// Initialize a keyring based on the provided options.
	kr, err := opts.Keyring()
	if err != nil {
		return nil, nil, err
	}

	// Sign the provided data using the key from the keyring.
	return kr.Sign(name, buf)
}
