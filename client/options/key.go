package options

import (
	"io"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptohd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
)

// Key represents options for key creation.
type Key struct {
	Account  uint32 `json:"account" toml:"account"`     // Account represents the account number.
	CoinType uint32 `json:"coin_type" toml:"coin_type"` // CoinType represents the coin type.
	Index    uint32 `json:"index" toml:"index"`         // Index represents the key index.
}

// NewKey creates a new Key instance with default values.
func NewKey() *Key {
	return &Key{
		CoinType: flags.DefaultKeyCoinType,
	}
}

// WithAccount sets the Account field and returns the modified Key instance.
func (k *Key) WithAccount(v uint32) *Key {
	k.Account = v
	return k
}

// WithCoinType sets the CoinType field and returns the modified Key instance.
func (k *Key) WithCoinType(v uint32) *Key {
	k.CoinType = v
	return k
}

// WithIndex sets the Index field and returns the modified Key instance.
func (k *Key) WithIndex(v uint32) *Key {
	k.Index = v
	return k
}

// HDPath returns the hierarchical deterministic (HD) path string based on CoinType, Account, and Index.
func (k *Key) HDPath() string {
	path := cryptohd.CreateHDPath(k.CoinType, k.Account, k.Index)
	return path.String()
}

// SignatureAlgo returns the default signature algorithm for keys.
func (k *Key) SignatureAlgo() keyring.SignatureAlgo {
	return cryptohd.Secp256k1
}

// AddKeyFlagsToCmd adds key-related flags to the given cobra command.
func AddKeyFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagKeyAccount(cmd)
	flags.SetFlagKeyCoinType(cmd)
	flags.SetFlagKeyIndex(cmd)
}

// NewKeyFromCmd creates and returns Key from the given cobra command's flags.
func NewKeyFromCmd(cmd *cobra.Command) (*Key, error) {
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

	// Return a new Key instance populated with the retrieved flag values.
	return &Key{
		Account:  account,
		CoinType: coinType,
		Index:    index,
	}, nil
}

// Keyring represents options for keyring creation.
type Keyring struct {
	AppName string    `json:"app_name" toml:"app_name"` // AppName is the name of the application.
	Backend string    `json:"backend" toml:"backend"`   // Backend is the keyring backend to use.
	HomeDir string    `json:"home_dir" toml:"home_dir"` // HomeDir is the directory to store keys.
	Input   io.Reader `json:"input" toml:"input"`       // Input is the input source for passphrase.
}

// NewKeyring creates a new Keyring instance with default values.
func NewKeyring() *Keyring {
	return &Keyring{
		AppName: flags.DefaultKeyringAppName,
		Backend: flags.DefaultKeyringBackend,
	}
}

// WithAppName sets the AppName field and returns the modified Keyring instance.
func (k *Keyring) WithAppName(v string) *Keyring {
	k.AppName = v
	return k
}

// WithBackend sets the Backend field and returns the modified Keyring instance.
func (k *Keyring) WithBackend(v string) *Keyring {
	k.Backend = v
	return k
}

// WithHomeDir sets the HomeDir field and returns the modified Keyring instance.
func (k *Keyring) WithHomeDir(v string) *Keyring {
	k.HomeDir = v
	return k
}

// WithInput sets the Input field and returns the modified Keyring instance.
func (k *Keyring) WithInput(v io.Reader) *Keyring {
	k.Input = v
	return k
}

// Keystore creates and returns a new keyring based on the provided options.
func (k *Keyring) Keystore(cdc codec.Codec) (keyring.Keyring, error) {
	return keyring.New(k.AppName, k.Backend, k.HomeDir, k.Input, cdc)
}

// AddKeyringFlagsToCmd adds keyring-related flags to the given cobra command.
func AddKeyringFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagKeyringAppName(cmd)
	flags.SetFlagKeyringBackend(cmd)
	flags.SetFlagKeyringHomeDir(cmd)
}

// NewKeyringFromCmd creates and returns Keyring from the given cobra command's flags.
func NewKeyringFromCmd(cmd *cobra.Command) (*Keyring, error) {
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

	// Return a new Keyring instance populated with the retrieved flag values.
	return &Keyring{
		AppName: appName,
		Backend: backend,
		HomeDir: homeDir,
		Input:   cmd.InOrStdin(),
	}, nil
}
