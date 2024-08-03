package options

import (
	"io"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptohd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
)

// KeyOptions represents options for key creation.
type KeyOptions struct {
	Account  uint32 `json:"account" toml:"account"`     // Account represents the account number.
	CoinType uint32 `json:"coin_type" toml:"coin_type"` // CoinType represents the coin type.
	Index    uint32 `json:"index" toml:"index"`         // Index represents the key index.
}

// NewDefaultKey creates a new KeyOptions instance with default values.
func NewDefaultKey() *KeyOptions {
	return &KeyOptions{
		CoinType: flags.DefaultKeyCoinType,
	}
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

// AddKeyFlagsToCmd adds key-related flags to the given cobra command.
func AddKeyFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagKeyAccount(cmd)
	flags.SetFlagKeyCoinType(cmd)
	flags.SetFlagKeyIndex(cmd)
}

// NewKeyOptionsFromCmd creates and returns KeyOptions from the given cobra command's flags.
func NewKeyOptionsFromCmd(cmd *cobra.Command) (*KeyOptions, error) {
	// Retrieve the account flag value from the command.
	account, err := flags.GetKeyAccountFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the coin type flag value from the command.
	coinType, err := flags.GetKeyCoinTypeFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the index flag value from the command.
	index, err := flags.GetKeyIndexFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new KeyOptions instance populated with the retrieved flag values.
	return &KeyOptions{
		Account:  account,
		CoinType: coinType,
		Index:    index,
	}, nil
}

// KeyringOptions represents options for keyring creation.
type KeyringOptions struct {
	AppName string    `json:"app_name" toml:"app_name"` // AppName is the name of the application.
	Backend string    `json:"backend" toml:"backend"`   // Backend is the keyring backend to use.
	HomeDir string    `json:"home_dir" toml:"home_dir"` // HomeDir is the directory to store keys.
	Input   io.Reader `json:"input" toml:"input"`       // Input is the input source for passphrase.
}

// NewDefaultKeyring creates a new KeyringOptions instance with default values.
func NewDefaultKeyring() *KeyringOptions {
	return &KeyringOptions{
		AppName: flags.DefaultKeyringAppName,
		Backend: flags.DefaultKeyringBackend,
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

// AddKeyringFlagsToCmd adds keyring-related flags to the given cobra command.
func AddKeyringFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagKeyringAppName(cmd)
	flags.SetFlagKeyringBackend(cmd)
	flags.SetFlagKeyringHomeDir(cmd)
}

// NewKeyringOptionsFromCmd creates and returns KeyringOptions from the given cobra command's flags.
func NewKeyringOptionsFromCmd(cmd *cobra.Command) (*KeyringOptions, error) {
	// Retrieve the application name flag value from the command.
	appName, err := flags.GetKeyringAppNameFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the backend flag value from the command.
	backend, err := flags.GetKeyringBackendFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the home directory flag value from the command.
	homeDir, err := flags.GetKeyringHomeDirFromCmd(cmd)
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
