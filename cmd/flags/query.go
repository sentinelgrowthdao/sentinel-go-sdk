package flags

import (
	"time"

	"github.com/spf13/cobra"
)

// Default values for page and query flags.
const (
	DefaultPageLimit       = 25
	DefaultQueryMaxRetries = 15
	DefaultQueryRPCAddr    = "https://rpc.sentinel.co:443"
	DefaultQueryTimeout    = 15 * time.Second
)

// GetPageCountTotalFromCmd retrieves the "page.count-total" flag value from the command.
func GetPageCountTotalFromCmd(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("page.count-total")
}

// GetPageKeyFromCmd retrieves the "page.key" flag value from the command.
func GetPageKeyFromCmd(cmd *cobra.Command) ([]byte, error) {
	return cmd.Flags().GetBytesBase64("page.key")
}

// GetPageLimitFromCmd retrieves the "page.limit" flag value from the command.
func GetPageLimitFromCmd(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("page.limit")
}

// GetPageOffsetFromCmd retrieves the "page.offset" flag value from the command.
func GetPageOffsetFromCmd(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("page.offset")
}

// GetPageReverseFromCmd retrieves the "page.reverse" flag value from the command.
func GetPageReverseFromCmd(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("page.reverse")
}

// SetFlagPageCountTotal adds the "page.count-total" flag to the command.
func SetFlagPageCountTotal(cmd *cobra.Command) {
	cmd.Flags().Bool("page.count-total", false, "Include total count in paged queries.")
}

// SetFlagPageKey adds the "page.key" flag to the command.
func SetFlagPageKey(cmd *cobra.Command) {
	cmd.Flags().BytesBase64("page.key", nil, "Base64-encoded key for page.")
}

// SetFlagPageLimit adds the "page.limit" flag to the command.
func SetFlagPageLimit(cmd *cobra.Command) {
	cmd.Flags().Uint64("page.limit", DefaultPageLimit, "Maximum number of results per page.")
}

// SetFlagPageOffset adds the "page.offset" flag to the command.
func SetFlagPageOffset(cmd *cobra.Command) {
	cmd.Flags().Uint64("page.offset", 0, "Offset for page.")
}

// SetFlagPageReverse adds the "page.reverse" flag to the command.
func SetFlagPageReverse(cmd *cobra.Command) {
	cmd.Flags().Bool("page.reverse", false, "Reverse the order of results in page.")
}

// GetQueryHeightFromCmd retrieves the "query.height" flag value from the command.
func GetQueryHeightFromCmd(cmd *cobra.Command) (int64, error) {
	return cmd.Flags().GetInt64("query.height")
}

// GetQueryMaxRetriesFromCmd retrieves the "query.max-retries" flag value from the command.
func GetQueryMaxRetriesFromCmd(cmd *cobra.Command) (int, error) {
	return cmd.Flags().GetInt("query.max-retries")
}

// GetQueryProveFromCmd retrieves the "query.prove" flag value from the command.
func GetQueryProveFromCmd(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("query.prove")
}

// GetQueryRPCAddrFromCmd retrieves the "query.rpc-addr" flag value from the command.
func GetQueryRPCAddrFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("query.rpc-addr")
}

// GetQueryTimeoutFromCmd retrieves the "query.timeout" flag value from the command.
func GetQueryTimeoutFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("query.timeout")
}

// SetFlagQueryHeight adds the "query.height" flag to the command.
func SetFlagQueryHeight(cmd *cobra.Command) {
	cmd.Flags().Int64("query.height", 0, "Block height at which the query is to be performed.")
}

// SetFlagQueryMaxRetries adds the "query.max-retries" flag to the command.
func SetFlagQueryMaxRetries(cmd *cobra.Command) {
	cmd.Flags().Int("query.max-retries", DefaultQueryMaxRetries, "Maximum number of retries for the query.")
}

// SetFlagQueryProve adds the "query.prove" flag to the command.
func SetFlagQueryProve(cmd *cobra.Command) {
	cmd.Flags().Bool("query.prove", false, "Include proof in query results.")
}

// SetFlagQueryRPCAddr adds the "query.rpc-addr" flag to the command.
func SetFlagQueryRPCAddr(cmd *cobra.Command) {
	cmd.Flags().String("query.rpc-addr", DefaultQueryRPCAddr, "Address of the RPC server.")
}

// SetFlagQueryTimeout adds the "query.timeout" flag to the command.
func SetFlagQueryTimeout(cmd *cobra.Command) {
	cmd.Flags().Duration("query.timeout", DefaultQueryTimeout, "Maximum duration for the query to be executed.")
}
