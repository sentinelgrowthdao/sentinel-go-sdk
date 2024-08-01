package options

import (
	"io"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptohd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/cobra"
)

// Default values for key and keyring options.
const (
	DefaultKeyCoinType    = 118
	DefaultKeyringAppName = "sentinel"
	DefaultKeyringBackend = "test"
)

// KeyOptions represents options for key creation.
type KeyOptions struct {
	Account  uint32 `json:"account,omitempty"`   // Account represents the account number.
	CoinType uint32 `json:"coin_type,omitempty"` // CoinType represents the coin type.
	Index    uint32 `json:"index,omitempty"`     // Index represents the key index.
}

// NewDefaultKeyOptions creates a new KeyOptions instance with default values.
func NewDefaultKeyOptions() *KeyOptions {
	return &KeyOptions{
		CoinType: DefaultKeyCoinType,
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

// GetKeyAccountFromCmd retrieves the "key.account" flag value from the command.
func GetKeyAccountFromCmd(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("key.account")
}

// GetKeyCoinTypeFromCmd retrieves the "key.coin-type" flag value from the command.
func GetKeyCoinTypeFromCmd(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("key.coin-type")
}

// GetKeyIndexFromCmd retrieves the "key.index" flag value from the command.
func GetKeyIndexFromCmd(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("key.index")
}

// SetFlagKeyAccount adds the "key.account" flag to the command.
func SetFlagKeyAccount(cmd *cobra.Command) {
	cmd.Flags().Uint32("key.account", 0, "Account number for key creation.")
}

// SetFlagKeyCoinType adds the "key.coin-type" flag to the command.
func SetFlagKeyCoinType(cmd *cobra.Command) {
	cmd.Flags().Uint32("key.coin-type", DefaultKeyCoinType, "Coin type for key creation.")
}

// SetFlagKeyIndex adds the "key.index" flag to the command.
func SetFlagKeyIndex(cmd *cobra.Command) {
	cmd.Flags().Uint32("key.index", 0, "Index for key creation.")
}

// AddKeyFlagsToCmd adds key-related flags to the given cobra command.
func AddKeyFlagsToCmd(cmd *cobra.Command) {
	SetFlagKeyAccount(cmd)
	SetFlagKeyCoinType(cmd)
	SetFlagKeyIndex(cmd)
}

// NewKeyOptionsFromCmd creates and returns KeyOptions from the given cobra command's flags.
func NewKeyOptionsFromCmd(cmd *cobra.Command) (*KeyOptions, error) {
	// Retrieve the value of the "key.account" flag.
	account, err := GetKeyAccountFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "key.coin-type" flag.
	coinType, err := GetKeyCoinTypeFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "key.index" flag.
	index, err := GetKeyIndexFromCmd(cmd)
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
	AppName string    `json:"app_name,omitempty"` // AppName is the name of the application.
	Backend string    `json:"backend,omitempty"`  // Backend is the keyring backend to use.
	HomeDir string    `json:"home_dir,omitempty"` // HomeDir is the directory to store keys.
	Input   io.Reader `json:"input,omitempty"`    // Input is the input source for passphrase.
}

// NewDefaultKeyringOptions creates a new KeyringOptions instance with default values.
func NewDefaultKeyringOptions() *KeyringOptions {
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

// GetKeyringAppNameFromCmd retrieves the "keyring.app-name" flag value from the command.
func GetKeyringAppNameFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("keyring.app-name")
}

// GetKeyringBackendFromCmd retrieves the "keyring.backend" flag value from the command.
func GetKeyringBackendFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("keyring.backend")
}

// GetKeyringHomeDirFromCmd retrieves the "keyring.home-dir" flag value from the command.
func GetKeyringHomeDirFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("keyring.home-dir")
}

// SetFlagKeyringAppName adds the "keyring.app-name" flag to the command.
func SetFlagKeyringAppName(cmd *cobra.Command) {
	cmd.Flags().String("keyring.app-name", DefaultKeyringAppName, "Name of the application.")
}

// SetFlagKeyringBackend adds the "keyring.backend" flag to the command.
func SetFlagKeyringBackend(cmd *cobra.Command) {
	cmd.Flags().String("keyring.backend", DefaultKeyringBackend, "Keyring backend to use.")
}

// SetFlagKeyringHomeDir adds the "keyring.home-dir" flag to the command.
func SetFlagKeyringHomeDir(cmd *cobra.Command) {
	cmd.Flags().String("keyring.home-dir", "", "Directory to store keys.")
}

// AddKeyringFlagsToCmd adds keyring-related flags to the given cobra command.
func AddKeyringFlagsToCmd(cmd *cobra.Command) {
	SetFlagKeyringAppName(cmd)
	SetFlagKeyringBackend(cmd)
	SetFlagKeyringHomeDir(cmd)
}

// NewKeyringOptionsFromCmd creates and returns KeyringOptions from the given cobra command's flags.
func NewKeyringOptionsFromCmd(cmd *cobra.Command) (*KeyringOptions, error) {
	// Retrieve the value of the "keyring.app-name" flag.
	appName, err := GetKeyringAppNameFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "keyring.backend" flag.
	backend, err := GetKeyringBackendFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "keyring.home-dir" flag.
	homeDir, err := GetKeyringHomeDirFromCmd(cmd)
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
