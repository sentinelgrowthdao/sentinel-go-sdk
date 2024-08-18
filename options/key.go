package options

import (
	cryptohd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/flags"
)

// Key represents options for key creation.
type Key struct {
	Account  uint32 `json:"account" toml:"account"`     // Account represents the account number in the key.
	CoinType uint32 `json:"coin_type" toml:"coin_type"` // CoinType represents the type of coin used.
	Index    uint32 `json:"index" toml:"index"`         // Index represents the specific key index.
}

// NewKey creates a new Key instance with default values.
func NewKey() *Key {
	return &Key{
		Account:  flags.DefaultKeyAccount,
		CoinType: flags.DefaultKeyCoinType,
		Index:    flags.DefaultKeyIndex,
	}
}

// WithAccount sets the Account field and returns the updated Key instance.
func (k *Key) WithAccount(v uint32) *Key {
	k.Account = v
	return k
}

// WithCoinType sets the CoinType field and returns the updated Key instance.
func (k *Key) WithCoinType(v uint32) *Key {
	k.CoinType = v
	return k
}

// WithIndex sets the Index field and returns the updated Key instance.
func (k *Key) WithIndex(v uint32) *Key {
	k.Index = v
	return k
}

// GetAccount returns the account number.
func (k *Key) GetAccount() uint32 {
	return k.Account
}

// GetCoinType returns the coin type.
func (k *Key) GetCoinType() uint32 {
	return k.CoinType
}

// GetIndex returns the key index.
func (k *Key) GetIndex() uint32 {
	return k.Index
}

// Validate validates all fields of the Key struct.
func (k *Key) Validate() error {
	return nil
}

// HDPath returns the hierarchical deterministic (HD) path string based on CoinType, Account, and Index.
func (k *Key) HDPath() string {
	v := cryptohd.CreateHDPath(k.GetCoinType(), k.GetAccount(), k.GetIndex())
	return v.String()
}

// SignatureAlgo returns the default signature algorithm for keys.
func (k *Key) SignatureAlgo() keyring.SignatureAlgo {
	return cryptohd.Secp256k1
}

// NewKeyFromCmd creates and returns a Key instance from the flags of the given cobra command.
func NewKeyFromCmd(cmd *cobra.Command) (*Key, error) {
	// Retrieve and validate the account flag value from the command.
	account, err := flags.GetKeyAccount(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve and validate the coin type flag value from the command.
	coinType, err := flags.GetKeyCoinType(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve and validate the index flag value from the command.
	index, err := flags.GetKeyIndex(cmd)
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
