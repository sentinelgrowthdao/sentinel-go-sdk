package flags

import (
	"github.com/spf13/cobra"
)

// Default values for query flags.
const (
	DefaultQueryHeight     = 0
	DefaultQueryMaxRetries = 15
	DefaultQueryProve      = false
	DefaultQueryRetryDelay = "1s"
	DefaultQueryRPCAddr    = "https://rpc.sentinel.co:443"
	DefaultQueryTimeout    = "15s"
)

// GetQueryHeight retrieves the "query.height" flag value from the command.
func GetQueryHeight(cmd *cobra.Command) (int64, error) {
	return cmd.Flags().GetInt64("query.height")
}

// GetQueryMaxRetries retrieves the "query.max-retries" flag value from the command.
func GetQueryMaxRetries(cmd *cobra.Command) (int, error) {
	return cmd.Flags().GetInt("query.max-retries")
}

// GetQueryProve retrieves the "query.prove" flag value from the command.
func GetQueryProve(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("query.prove")
}

// GetQueryRetryDelay retrieves the "query.retry-delay" flag value from the command.
func GetQueryRetryDelay(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("query.retry-delay")
}

// GetQueryRPCAddr retrieves the "query.rpc-addr" flag value from the command.
func GetQueryRPCAddr(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("query.rpc-addr")
}

// GetQueryTimeout retrieves the "query.timeout" flag value from the command.
func GetQueryTimeout(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("query.timeout")
}

// SetFlagQueryHeight adds the "query.height" flag to the command.
func SetFlagQueryHeight(cmd *cobra.Command) {
	cmd.Flags().Int64("query.height", DefaultQueryHeight, "Block height at which the query is to be performed.")
}

// SetFlagQueryMaxRetries adds the "query.max-retries" flag to the command.
func SetFlagQueryMaxRetries(cmd *cobra.Command) {
	cmd.Flags().Int("query.max-retries", DefaultQueryMaxRetries, "Maximum number of retries for the query.")
}

// SetFlagQueryProve adds the "query.prove" flag to the command.
func SetFlagQueryProve(cmd *cobra.Command) {
	cmd.Flags().Bool("query.prove", DefaultQueryProve, "Include proof in query results.")
}

// SetFlagQueryRetryDelay adds the "query.retry-delay" flag to the command.
func SetFlagQueryRetryDelay(cmd *cobra.Command) {
	cmd.Flags().String("query.retry-delay", DefaultQueryRetryDelay, "Delay between retries for the query.")
}

// SetFlagQueryRPCAddr adds the "query.rpc-addr" flag to the command.
func SetFlagQueryRPCAddr(cmd *cobra.Command) {
	cmd.Flags().String("query.rpc-addr", DefaultQueryRPCAddr, "Address of the RPC server.")
}

// SetFlagQueryTimeout adds the "query.timeout" flag to the command.
func SetFlagQueryTimeout(cmd *cobra.Command) {
	cmd.Flags().String("query.timeout", DefaultQueryTimeout, "Maximum duration for the query to be executed.")
}

// AddQueryFlags adds query-related flags to the given cobra command.
func AddQueryFlags(cmd *cobra.Command) {
	SetFlagQueryHeight(cmd)
	SetFlagQueryMaxRetries(cmd)
	SetFlagQueryProve(cmd)
	SetFlagQueryRetryDelay(cmd)
	SetFlagQueryRPCAddr(cmd)
	SetFlagQueryTimeout(cmd)
}
