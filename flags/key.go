package flags

import (
	"github.com/spf13/cobra"
)

// Default values for key flags.
const (
	DefaultKeyAccount  = 0
	DefaultKeyCoinType = 118
	DefaultKeyIndex    = 0
)

// GetKeyAccount retrieves the "key.account" flag value from the command.
func GetKeyAccount(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("key.account")
}

// GetKeyCoinType retrieves the "key.coin-type" flag value from the command.
func GetKeyCoinType(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("key.coin-type")
}

// GetKeyIndex retrieves the "key.index" flag value from the command.
func GetKeyIndex(cmd *cobra.Command) (uint32, error) {
	return cmd.Flags().GetUint32("key.index")
}

// SetFlagKeyAccount adds the "key.account" flag to the command.
func SetFlagKeyAccount(cmd *cobra.Command) {
	cmd.Flags().Uint32("key.account", DefaultKeyAccount, "Account number for key creation.")
}

// SetFlagKeyCoinType adds the "key.coin-type" flag to the command.
func SetFlagKeyCoinType(cmd *cobra.Command) {
	cmd.Flags().Uint32("key.coin-type", DefaultKeyCoinType, "Coin type for key creation.")
}

// SetFlagKeyIndex adds the "key.index" flag to the command.
func SetFlagKeyIndex(cmd *cobra.Command) {
	cmd.Flags().Uint32("key.index", DefaultKeyIndex, "Index for key creation.")
}

// AddKeyFlags adds key-related flags to the given cobra command.
func AddKeyFlags(cmd *cobra.Command) {
	SetFlagKeyAccount(cmd)
	SetFlagKeyCoinType(cmd)
	SetFlagKeyIndex(cmd)
}
