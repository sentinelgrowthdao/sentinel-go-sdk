package options

import (
	"github.com/cometbft/cometbft/config"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
)

// Log holds the configuration options for logging.
type Log struct {
	Format string `json:"format" toml:"format"` // Log format (e.g., plain or JSON)
	Level  string `json:"level" toml:"level"`   // Log level (e.g., info, debug, error)
}

// NewLog initializes a Log instance with default values.
func NewLog() *Log {
	return &Log{
		Format: config.LogFormatPlain,
		Level:  config.DefaultLogLevel,
	}
}

// WithFormat sets the log format and returns the updated Log instance.
func (k *Log) WithFormat(v string) *Log {
	k.Format = v
	return k
}

// WithLevel sets the log level and returns the updated Log instance.
func (k *Log) WithLevel(v string) *Log {
	k.Level = v
	return k
}

// AddLogFlagsToCmd attaches logging-related flags to the provided cobra command.
func AddLogFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagLogFormat(cmd)
	flags.SetFlagLogLevel(cmd)
}

// NewLogFromCmd extracts and returns Log from the given cobra command's flags.
func NewLogFromCmd(cmd *cobra.Command) (*Log, error) {
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

	// Create and return a Log instance with the retrieved values.
	return &Log{
		Format: format,
		Level:  level,
	}, nil
}
