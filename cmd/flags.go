package cmd

import (
	"github.com/spf13/cobra"
)

const (
	DefaultOutputFormat = "text"
)

// SetFlagOutputFormat adds a flag for specifying the output format to the given command.
func SetFlagOutputFormat(cmd *cobra.Command) {
	cmd.Flags().String("output-format", DefaultOutputFormat, "Specify the output format (json or text)")
}

// GetOutputFormatFromCmd retrieves the output format flag value from the given command.
func GetOutputFormatFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("output-format")
}
