package flags

import (
	"github.com/spf13/cobra"
)

// Default values for keyring flags.
const (
	DefaultKeyringAppName = "sentinel"
	DefaultKeyringBackend = "test"
	DefaultKeyringHomeDir = ""
)

// GetKeyringAppName retrieves the "keyring.app-name" flag value from the command.
func GetKeyringAppName(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("keyring.app-name")
}

// GetKeyringBackend retrieves the "keyring.backend" flag value from the command.
func GetKeyringBackend(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("keyring.backend")
}

// GetKeyringHomeDir retrieves the "keyring.home-dir" flag value from the command.
func GetKeyringHomeDir(cmd *cobra.Command) (string, error) {
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
	cmd.Flags().String("keyring.home-dir", DefaultKeyringHomeDir, "Directory to store keys.")
}

// AddKeyringFlags adds keyring-related flags to the given cobra command.
func AddKeyringFlags(cmd *cobra.Command) {
	SetFlagKeyringAppName(cmd)
	SetFlagKeyringBackend(cmd)
	SetFlagKeyringHomeDir(cmd)
}
