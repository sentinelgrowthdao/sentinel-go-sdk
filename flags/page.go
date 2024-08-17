package flags

import (
	"github.com/spf13/cobra"
)

// Default values for pagination flags.
const (
	DefaultPageCountTotal = false
	DefaultPageKey        = ""
	DefaultPageLimit      = 25
	DefaultPageOffset     = 0
	DefaultPageReverse    = false
)

// GetPageCountTotal retrieves the "page.count-total" flag value from the command.
func GetPageCountTotal(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("page.count-total")
}

// GetPageKey retrieves the "page.key" flag value from the command.
func GetPageKey(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("page.key")
}

// GetPageLimit retrieves the "page.limit" flag value from the command.
func GetPageLimit(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("page.limit")
}

// GetPageOffset retrieves the "page.offset" flag value from the command.
func GetPageOffset(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("page.offset")
}

// GetPageReverse retrieves the "page.reverse" flag value from the command.
func GetPageReverse(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("page.reverse")
}

// SetFlagPageCountTotal adds the "page.count-total" flag to the command.
func SetFlagPageCountTotal(cmd *cobra.Command) {
	cmd.Flags().Bool("page.count-total", DefaultPageCountTotal, "Include total count in paged queries.")
}

// SetFlagPageKey adds the "page.key" flag to the command.
func SetFlagPageKey(cmd *cobra.Command) {
	cmd.Flags().String("page.key", DefaultPageKey, "Base64-encoded key for page.")
}

// SetFlagPageLimit adds the "page.limit" flag to the command.
func SetFlagPageLimit(cmd *cobra.Command) {
	cmd.Flags().Uint64("page.limit", DefaultPageLimit, "Maximum number of results per page.")
}

// SetFlagPageOffset adds the "page.offset" flag to the command.
func SetFlagPageOffset(cmd *cobra.Command) {
	cmd.Flags().Uint64("page.offset", DefaultPageOffset, "Offset for page.")
}

// SetFlagPageReverse adds the "page.reverse" flag to the command.
func SetFlagPageReverse(cmd *cobra.Command) {
	cmd.Flags().Bool("page.reverse", DefaultPageReverse, "Reverse the order of results in page.")
}

// AddPageFlags adds page-related flags to the given cobra command.
func AddPageFlags(cmd *cobra.Command) {
	SetFlagPageCountTotal(cmd)
	SetFlagPageKey(cmd)
	SetFlagPageLimit(cmd)
	SetFlagPageOffset(cmd)
	SetFlagPageReverse(cmd)
}
