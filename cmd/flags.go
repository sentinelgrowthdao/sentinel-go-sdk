package cmd

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/spf13/cobra"
)

const (
	flagOutputFormat = "output-format"
)

// SetFlagOutputFormat adds a flag for specifying the output format to the given command.
func SetFlagOutputFormat(cmd *cobra.Command) {
	cmd.Flags().String(flagOutputFormat, keys.OutputFormatText, "Specify the output format (json or text)")
}

// GetOutputFormatFromCmd retrieves the output format flag value from the given command and validates it.
func GetOutputFormatFromCmd(cmd *cobra.Command) (string, error) {
	s, err := cmd.Flags().GetString(flagOutputFormat)
	if err != nil {
		return "", err
	}

	if s != keys.OutputFormatJSON && s != keys.OutputFormatText {
		return "", fmt.Errorf("invalid output format: %s", s)
	}

	return s, nil
}
