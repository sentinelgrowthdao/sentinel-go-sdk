package flags

import (
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
