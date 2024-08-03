package flags

import (
	"github.com/spf13/cobra"
)

// Default values for key and keyring flags.
const (
	DefaultKeyCoinType    = 118
	DefaultKeyringAppName = "sentinel"
	DefaultKeyringBackend = "test"
)

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
