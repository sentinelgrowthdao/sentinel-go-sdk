package options

import (
	"errors"

	"github.com/cometbft/cometbft/config"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/flags"
)

// Log holds the configuration options for logging.
type Log struct {
	Format string `json:"format" toml:"format"` // Log format (e.g., plain or JSON).
	Level  string `json:"level" toml:"level"`   // Log level (e.g., info, debug, error).
}

// NewLog initializes a Log instance with default values.
func NewLog() *Log {
	return &Log{
		Format: config.LogFormatPlain,
		Level:  config.DefaultLogLevel,
	}
}

// WithFormat sets the log format and returns the updated Log instance.
func (l *Log) WithFormat(v string) *Log {
	l.Format = v
	return l
}

// WithLevel sets the log level and returns the updated Log instance.
func (l *Log) WithLevel(v string) *Log {
	l.Level = v
	return l
}

// GetFormat returns the log format.
func (l *Log) GetFormat() string {
	return l.Format
}

// GetLevel returns the log level.
func (l *Log) GetLevel() string {
	return l.Level
}

// ValidateLogFormat checks if the Format field is valid.
func ValidateLogFormat(v string) error {
	allowedFormats := map[string]bool{
		"json":  true,
		"plain": true,
	}

	if v == "" {
		return errors.New("format must be non-empty")
	}
	if _, ok := allowedFormats[v]; !ok {
		return errors.New("format must be one of: json, plain")
	}

	return nil
}

// ValidateLogLevel checks if the Level field is valid.
func ValidateLogLevel(v string) error {
	validLevels := map[string]bool{
		"debug": true,
		"error": true,
		"fatal": true,
		"info":  true,
		"panic": true,
		"warn":  true,
	}

	if v == "" {
		return errors.New("level must be non-empty")
	}
	if _, ok := validLevels[v]; !ok {
		return errors.New("level must be one of: debug, error, fatal, info, panic, warn")
	}

	return nil
}

// Validate checks all fields of the Log struct for validity.
func (l *Log) Validate() error {
	if err := ValidateLogFormat(l.Format); err != nil {
		return err
	}
	if err := ValidateLogLevel(l.Level); err != nil {
		return err
	}

	return nil
}

// NewLogFromCmd extracts and returns a Log instance from the given cobra command's flags.
func NewLogFromCmd(cmd *cobra.Command) (*Log, error) {
	// Retrieve the log format value from the command's flags.
	format, err := flags.GetLogFormat(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the log level value from the command's flags.
	level, err := flags.GetLogLevel(cmd)
	if err != nil {
		return nil, err
	}

	// Create and return a Log instance with the retrieved values.
	return &Log{
		Format: format,
		Level:  level,
	}, nil
}
