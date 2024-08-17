package client

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/go-bip39"
)

// Key retrieves key information from the keyring based on the provided name and options.
// It initializes a keyring using the provided options and returns the key information.
func (c *Client) Key(name string, opts *Options) (*keyring.Record, error) {
	// Initialize a keyring based on the provided options.
	kr, err := c.Keyring(opts)
	if err != nil {
		return nil, err
	}

	// Retrieve key information from the keyring.
	return kr.Key(name)
}

// Sign signs the provided data using the key from the keyring specified by the name and options.
// It initializes a keyring, retrieves the key, and signs the data.
func (c *Client) Sign(name string, buf []byte, opts *Options) ([]byte, cryptotypes.PubKey, error) {
	// Initialize a keyring based on the provided options.
	kr, err := c.Keyring(opts)
	if err != nil {
		return nil, nil, err
	}

	// Sign the provided data using the key from the keyring.
	return kr.Sign(name, buf)
}

// Keys retrieves a list of all keys from the keyring based on the provided options.
// It initializes a keyring and returns a list of key records.
func (c *Client) Keys(opts *Options) ([]*keyring.Record, error) {
	// Initialize a keyring based on the provided options.
	kr, err := c.Keyring(opts)
	if err != nil {
		return nil, err
	}

	// Retrieve and return the list of key records.
	return kr.List()
}

// DeleteKey deletes the key from the keyring based on the provided name and options.
// It initializes a keyring and removes the key specified by the name.
func (c *Client) DeleteKey(name string, opts *Options) error {
	// Initialize a keyring based on the provided options.
	kr, err := c.Keyring(opts)
	if err != nil {
		return err
	}

	// Delete the key from the keyring.
	return kr.Delete(name)
}

// NewMnemonic generates a new mnemonic phrase using bip39 with 256 bits of entropy.
// It returns the generated mnemonic or an error if one occurs.
func (c *Client) NewMnemonic() (string, error) {
	// Generate new entropy for the mnemonic.
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	// Create a new mnemonic phrase from the entropy.
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

// CreateKey creates a new key in the keyring with the provided name, mnemonic, and bip39 passphrase.
// If mnemonic is empty, a new mnemonic is generated. It returns the mnemonic, the created key record, or an error.
func (c *Client) CreateKey(name, mnemonic, bip39Pass string, opts *Options) (string, *keyring.Record, error) {
	// Initialize a keyring based on the provided options.
	kr, err := c.Keyring(opts)
	if err != nil {
		return "", nil, err
	}

	// Generate a new mnemonic if none is provided.
	if mnemonic == "" {
		mnemonic, err = c.NewMnemonic()
		if err != nil {
			return "", nil, err
		}
	}

	// Create a new key with the provided or newly generated mnemonic.
	key, err := kr.NewAccount(name, mnemonic, bip39Pass, opts.HDPath(), opts.SignatureAlgo())
	if err != nil {
		return "", nil, err
	}

	return mnemonic, key, nil
}
