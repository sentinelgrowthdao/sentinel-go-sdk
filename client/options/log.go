package options

import (
	"github.com/cometbft/cometbft/config"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
)

// LogOptions holds the configuration options for logging.
type LogOptions struct {
	Format string `json:"format" toml:"format"` // Log format (e.g., plain or JSON)
	Level  string `json:"level" toml:"level"`   // Log level (e.g., info, debug, error)
}

// NewDefaultLog initializes a LogOptions instance with default values.
func NewDefaultLog() *LogOptions {
	return &LogOptions{
		Format: config.LogFormatPlain,
		Level:  config.DefaultLogLevel,
	}
}

// WithFormat sets the log format and returns the updated LogOptions instance.
func (k *LogOptions) WithFormat(v string) *LogOptions {
	k.Format = v
	return k
}

// WithLevel sets the log level and returns the updated LogOptions instance.
func (k *LogOptions) WithLevel(v string) *LogOptions {
	k.Level = v
	return k
}

// AddLogFlagsToCmd attaches logging-related flags to the provided cobra command.
func AddLogFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagLogFormat(cmd)
	flags.SetFlagLogLevel(cmd)
}

// NewLogOptionsFromCmd extracts and returns LogOptions from the given cobra command's flags.
func NewLogOptionsFromCmd(cmd *cobra.Command) (*LogOptions, error) {
	// Retrieve the log format value from the command's flags.
	format, err := flags.GetLogFormatFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the log level value from the command's flags.
	level, err := flags.GetLogLevelFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Create and return a LogOptions instance with the retrieved values.
	return &LogOptions{
		Format: format,
		Level:  level,
	}, nil
}
