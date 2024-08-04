package log

import (
	"io"

	"cosmossdk.io/log"
	"github.com/cometbft/cometbft/config"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
)

// NewLogger creates a new logger instance with the specified output writer, format, and log level.
func NewLogger(w io.Writer, format, level string) (log.Logger, error) {
	// Parse the log level from the string
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	// Prepare options for logger
	options := []log.Option{
		log.LevelOption(logLevel),
	}

	// Set log format based on the provided format string
	if format == config.LogFormatJSON {
		options = append(options, log.OutputJSONOption())
	}

	// Create and return the logger with the specified options
	return log.NewLogger(w, options...), nil
}

// NewLoggerFromCmd creates a new logger instance based on command-line flags from the provided Cobra command.
func NewLoggerFromCmd(cmd *cobra.Command) (log.Logger, error) {
	// Retrieve log format and level from command-line flags
	format, err := flags.GetLogFormatFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	level, err := flags.GetLogLevelFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Create and return the logger with the retrieved format and level
	return NewLogger(cmd.OutOrStderr(), format, level)
}
