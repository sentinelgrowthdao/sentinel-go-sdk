package flags

import (
	"github.com/cometbft/cometbft/config"
	"github.com/spf13/cobra"
)

// GetLogFormat retrieves the log format flag value from the command.
func GetLogFormat(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("log.format")
}

// GetLogLevel retrieves the log level flag value from the command.
func GetLogLevel(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("log.level")
}

// SetFlagLogFormat adds a flag for specifying the log format to the given command.
func SetFlagLogFormat(cmd *cobra.Command) {
	cmd.Flags().String("log.format", config.LogFormatPlain, "Specify the log format (json or plain).")
}

// SetFlagLogLevel adds a flag for specifying the log level to the given command.
func SetFlagLogLevel(cmd *cobra.Command) {
	cmd.Flags().String("log.level", config.DefaultLogLevel, "Specify the log level (debug, info, warn, error, fatal, panic).")
}

// AddLogFlags attaches logging-related flags to the provided cobra command.
func AddLogFlags(cmd *cobra.Command) {
	SetFlagLogFormat(cmd)
	SetFlagLogLevel(cmd)
}
