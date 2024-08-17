package log

import (
	"io"

	"cosmossdk.io/log"
	"github.com/cometbft/cometbft/config"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/options"
)

// NewLogger creates a new logger instance with the specified output writer, format, and log level.
func NewLogger(w io.Writer, format, level string) (log.Logger, error) {
	// Parse the log level from the string
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	// Prepare options for logger
	opts := []log.Option{
		log.LevelOption(logLevel),
	}

	// Set log format based on the provided format string
	if format == config.LogFormatJSON {
		opts = append(opts, log.OutputJSONOption())
	}

	// Create and return the logger with the specified options
	return log.NewLogger(w, opts...), nil
}

// NewLoggerFromCmd creates a new logger instance based on command-line flags from the provided Cobra command.
func NewLoggerFromCmd(cmd *cobra.Command) (log.Logger, error) {
	// Retrieve log format and level from command-line flags
	opts, err := options.NewLogFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Create and return the logger with the retrieved format and level
	return NewLogger(cmd.OutOrStderr(), opts.GetFormat(), opts.GetLevel())
}
