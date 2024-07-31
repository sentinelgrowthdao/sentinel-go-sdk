package options

import (
	"io"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptohd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/cobra"
)

// Default values for keyring options.
const (
	DefaultKeyringAppName = "sentinel"
	DefaultKeyringBackend = "test"
)

// KeyringOptions represents options for keyring creation.
type KeyringOptions struct {
	AppName string    `json:"app_name,omitempty"` // AppName is the name of the application.
	Backend string    `json:"backend,omitempty"`  // Backend is the keyring backend to use.
	HomeDir string    `json:"home_dir,omitempty"` // HomeDir is the directory to store keys.
	Input   io.Reader `json:"input,omitempty"`    // Input is the input source for passphrase.
}

// Keyring creates a new KeyringOptions instance with default values.
func Keyring() *KeyringOptions {
	return &KeyringOptions{
		AppName: DefaultKeyringAppName,
		Backend: DefaultKeyringBackend,
	}
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

// Keyring creates and returns a new keyring based on the provided options.
func (k *KeyringOptions) Keyring(cdc codec.Codec) (keyring.Keyring, error) {
	return keyring.New(k.AppName, k.Backend, k.HomeDir, k.Input, cdc)
}

// AddKeyringFlagsToCmd adds keyring related flags to the given cobra command.
func AddKeyringFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().String("keyring.app-name", DefaultKeyringAppName, "Application name for keyring isolation.")
	cmd.Flags().String("keyring.backend", DefaultKeyringBackend, "Type of keyring backend to use (e.g., file, os, pass).")
	cmd.Flags().String("keyring.home-dir", "", "Directory path for storing keyring data.")
}

// NewKeyringOptionsFromCmd creates and returns KeyringOptions from the given cobra command's flags.
func NewKeyringOptionsFromCmd(cmd *cobra.Command) (*KeyringOptions, error) {
	// Retrieve the value of the "keyring.app-name" flag.
	appName, err := cmd.Flags().GetString("keyring.app-name")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "keyring.backend" flag.
	backend, err := cmd.Flags().GetString("keyring.backend")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "keyring.home-dir" flag.
	homeDir, err := cmd.Flags().GetString("keyring.home-dir")
	if err != nil {
		return nil, err
	}

	// Return a new KeyringOptions instance populated with the retrieved flag values.
	return &KeyringOptions{
		AppName: appName,
		Backend: backend,
		HomeDir: homeDir,
		Input:   cmd.InOrStdin(),
	}, nil
}

// Default values for key options.
const (
	DefaultKeyAccount  = 0
	DefaultKeyCoinType = 118
	DefaultKeyIndex    = 0
)

// KeyOptions represents options for key creation.
type KeyOptions struct {
	*KeyringOptions        // Embedding KeyringOptions for composition.
	Account         uint32 `json:"account,omitempty"`   // Account represents the account number.
	CoinType        uint32 `json:"coin_type,omitempty"` // CoinType represents the coin type.
	Index           uint32 `json:"index,omitempty"`     // Index represents the key index.
}

// Key creates a new KeyOptions instance with default values.
func Key() *KeyOptions {
	return &KeyOptions{
		KeyringOptions: Keyring(), // Initialize embedded KeyringOptions.
		Account:        DefaultKeyAccount,
		CoinType:       DefaultKeyCoinType,
		Index:          DefaultKeyIndex,
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

// WithCoinType sets the CoinType field and returns the modified KeyOptions instance.
func (k *KeyOptions) WithCoinType(v uint32) *KeyOptions {
	k.CoinType = v
	return k
}

// WithIndex sets the Index field and returns the modified KeyOptions instance.
func (k *KeyOptions) WithIndex(v uint32) *KeyOptions {
	k.Index = v
	return k
}

// HDPath returns the hierarchical deterministic (HD) path string based on CoinType, Account, and Index.
func (k *KeyOptions) HDPath() string {
	path := cryptohd.CreateHDPath(k.CoinType, k.Account, k.Index)
	return path.String()
}

// SignatureAlgo returns the default signature algorithm for keys.
func (k *KeyOptions) SignatureAlgo() keyring.SignatureAlgo {
	return cryptohd.Secp256k1
}

// AddKeyFlagsToCmd adds keyring and key-related flags to the given cobra command.
func AddKeyFlagsToCmd(cmd *cobra.Command) {
	// Add keyring-related flags to the command.
	AddKeyringFlagsToCmd(cmd)

	cmd.Flags().Uint32("key.account", DefaultKeyAccount, "Specifies the account number in the HD wallet's hierarchical structure.")
	cmd.Flags().Uint32("key.coin-type", DefaultKeyCoinType, "Indicates the type or purpose of the key within the HD wallet structure.")
	cmd.Flags().Uint32("key.index", DefaultKeyIndex, "Designates the key index within the specified HD wallet account.")
}

// NewKeyOptionsFromCmd creates and returns KeyOptions from the given cobra command's flags.
func NewKeyOptionsFromCmd(cmd *cobra.Command) (*KeyOptions, error) {
	// Retrieve and create KeyringOptions from the command's flags.
	keyringOpts, err := NewKeyringOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "key.account" flag.
	account, err := cmd.Flags().GetUint32("key.account")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "key.coin-type" flag.
	coinType, err := cmd.Flags().GetUint32("key.coin-type")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "key.index" flag.
	index, err := cmd.Flags().GetUint32("key.index")
	if err != nil {
		return nil, err
	}

	// Return a new KeyOptions instance populated with the retrieved flag values and KeyringOptions.
	return &KeyOptions{
		KeyringOptions: keyringOpts,
		Account:        account,
		CoinType:       coinType,
		Index:          index,
	}, nil
}
