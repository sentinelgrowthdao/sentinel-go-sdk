package options

import (
	"io"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

// DefaultKeyType represents the default key type.
const DefaultKeyType = 118

// KeyringOptions represents options for keyring creation.
type KeyringOptions struct {
	AppName string    `json:"app_name,omitempty"` // AppName is the name of the application.
	Backend string    `json:"backend,omitempty"`  // Backend is the keyring backend to use.
	HomeDir string    `json:"home_dir,omitempty"` // HomeDir is the directory to store keys.
	Input   io.Reader `json:"input,omitempty"`    // Input is the input source for passphrase.
}

// Keyring creates a new KeyringOptions instance.
func Keyring() *KeyringOptions {
	return &KeyringOptions{}
}

// WithAppName sets the AppName field and returns the modified KeyringOptions instance.
func (k *KeyringOptions) WithAppName(v string) *KeyringOptions {
	k.AppName = v
	return k
}

// WithBackend sets the Backend field and returns the modified KeyringOptions instance.
func (k *KeyringOptions) WithBackend(v string) *KeyringOptions {
	k.Backend = v
	return k
}

// WithHomeDir sets the HomeDir field and returns the modified KeyringOptions instance.
func (k *KeyringOptions) WithHomeDir(v string) *KeyringOptions {
	k.HomeDir = v
	return k
}

// WithInput sets the Input field and returns the modified KeyringOptions instance.
func (k *KeyringOptions) WithInput(v io.Reader) *KeyringOptions {
	k.Input = v
	return k
}

// Keyring returns a new keyring based on the provided options.
func (k *KeyringOptions) Keyring() (keyring.Keyring, error) {
	return keyring.New(k.AppName, k.Backend, k.HomeDir, k.Input)
}

// KeyOptions represents options for key creation.
type KeyOptions struct {
	*KeyringOptions        // Embedding KeyringOptions for composition.
	Account         uint32 `json:"account,omitempty"` // Account represents the account number.
	Index           uint32 `json:"index,omitempty"`   // Index represents the key index.
	Type            uint32 `json:"type,omitempty"`    // Type represents the key type.
}

// Key creates a new KeyOptions instance with default values.
func Key() *KeyOptions {
	return &KeyOptions{
		KeyringOptions: Keyring(), // Initialize embedded KeyringOptions.
		Type:           DefaultKeyType,
	}
}

// WithKeyringOptions sets the KeyringOptions field and returns the modified KeyOptions instance.
func (k *KeyOptions) WithKeyringOptions(v *KeyringOptions) *KeyOptions {
	k.KeyringOptions = v
	return k
}

// WithAccount sets the Account field and returns the modified KeyOptions instance.
func (k *KeyOptions) WithAccount(v uint32) *KeyOptions {
	k.Account = v
	return k
}

// WithIndex sets the Index field and returns the modified KeyOptions instance.
func (k *KeyOptions) WithIndex(v uint32) *KeyOptions {
	k.Index = v
	return k
}

// WithType sets the Type field and returns the modified KeyOptions instance.
func (k *KeyOptions) WithType(v uint32) *KeyOptions {
	k.Type = v
	return k
}
