package flags

import (
	"github.com/cometbft/cometbft/config"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/spf13/cobra"
)

// SetFlagOutputFormat adds a flag for specifying the output format to the given command.
func SetFlagOutputFormat(cmd *cobra.Command) {
	cmd.Flags().String("output-format", keys.OutputFormatText, "Specify the output format (json or text)")
}

// GetOutputFormatFromCmd retrieves the output format flag value from the given command.
func GetOutputFormatFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("output-format")
}

// SetFlagLogFormat adds a flag for specifying the log format to the given command.
func SetFlagLogFormat(cmd *cobra.Command) {
	cmd.Flags().String("log-format", config.LogFormatJSON, "Specify the log format (json or plain)")
}

// GetLogFormatFromCmd retrieves the log format flag value from the given command.
func GetLogFormatFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("log-format")
}

// SetFlagLogLevel adds a flag for specifying the log level to the given command.
func SetFlagLogLevel(cmd *cobra.Command) {
	cmd.Flags().String("log-level", config.DefaultLogLevel, "Specify the log level (debug, info, warn, error, fatal, panic)")
}

// GetLogLevelFromCmd retrieves the log level flag value from the given command.
func GetLogLevelFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("log-level")
}
