package options

import (
	"errors"
	"io"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/flags"
)

// Keyring represents options for keyring creation.
type Keyring struct {
	Input io.Reader // Input is the source of passphrase input.

	AppName string `json:"app_name" toml:"app_name"` // AppName is the name of the application.
	Backend string `json:"backend" toml:"backend"`   // Backend specifies the keyring backend to use.
	HomeDir string `json:"home_dir" toml:"home_dir"` // HomeDir is the directory where keys are stored.
}

// NewKeyring creates a new Keyring instance with default values.
func NewKeyring() *Keyring {
	return &Keyring{
		AppName: flags.DefaultKeyringAppName,
		Backend: flags.DefaultKeyringBackend,
		HomeDir: flags.DefaultKeyringHomeDir,
	}
}

// WithAppName sets the AppName field and returns the updated Keyring instance.
func (k *Keyring) WithAppName(v string) *Keyring {
	k.AppName = v
	return k
}

// WithBackend sets the Backend field and returns the updated Keyring instance.
func (k *Keyring) WithBackend(v string) *Keyring {
	k.Backend = v
	return k
}

// WithHomeDir sets the HomeDir field and returns the updated Keyring instance.
func (k *Keyring) WithHomeDir(v string) *Keyring {
	k.HomeDir = v
	return k
}

// WithInput sets the Input field and returns the updated Keyring instance.
func (k *Keyring) WithInput(v io.Reader) *Keyring {
	k.Input = v
	return k
}

// GetAppName returns the application name.
func (k *Keyring) GetAppName() string {
	return k.AppName
}

// GetBackend returns the keyring backend.
func (k *Keyring) GetBackend() string {
	return k.Backend
}

// GetHomeDir returns the home directory for storing keys.
func (k *Keyring) GetHomeDir() string {
	return k.HomeDir
}

// GetInput returns the input source for the passphrase.
func (k *Keyring) GetInput() io.Reader {
	return k.Input
}

// ValidateKeyringAppName checks if the AppName field is valid.
func ValidateKeyringAppName(v string) error {
	if v == "" {
		return errors.New("app name must be non-empty")
	}

	return nil
}

// ValidateKeyringBackend checks if the Backend field is valid.
func ValidateKeyringBackend(v string) error {
	allowedBackends := map[string]bool{
		"file":    true,
		"kwallet": true,
		"memory":  true,
		"os":      true,
		"pass":    true,
		"test":    true,
	}

	if v == "" {
		return errors.New("backend must be non-empty")
	}
	if _, ok := allowedBackends[v]; !ok {
		return errors.New("backend must be one of: file, kwallet, memory, os, pass, test")
	}

	return nil
}

// ValidateKeyringHomeDir checks if the HomeDir field is valid.
func ValidateKeyringHomeDir(v string) error {
	if v == "" {
		return errors.New("home directory must be non-empty")
	}

	return nil
}

// Validate validates all fields of the Keyring struct.
func (k *Keyring) Validate() error {
	if err := ValidateKeyringAppName(k.AppName); err != nil {
		return err
	}
	if err := ValidateKeyringBackend(k.Backend); err != nil {
		return err
	}
	if err := ValidateKeyringHomeDir(k.HomeDir); err != nil {
		return err
	}

	return nil
}

// Keystore creates and returns a new keyring based on the provided options.
func (k *Keyring) Keystore(cdc codec.Codec) (keyring.Keyring, error) {
	return keyring.New(k.GetAppName(), k.GetBackend(), k.GetHomeDir(), k.GetInput(), cdc)
}

// NewKeyringFromCmd creates and returns a Keyring from the given cobra command's flags.
func NewKeyringFromCmd(cmd *cobra.Command) (*Keyring, error) {
	// Retrieve the application name flag value from the command.
	appName, err := flags.GetKeyringAppName(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the backend flag value from the command.
	backend, err := flags.GetKeyringBackend(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the home directory flag value from the command.
	homeDir, err := flags.GetKeyringHomeDir(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Keyring instance populated with the retrieved flag values.
	return &Keyring{
		AppName: appName,
		Backend: backend,
		HomeDir: homeDir,
		Input:   cmd.InOrStdin(), // Use the command's input or standard input as the passphrase source.
	}, nil
}
