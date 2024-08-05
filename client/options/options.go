package options

import (
	"github.com/spf13/cobra"
)

// Options aggregates all the individual option structs for a comprehensive configuration.
type Options struct {
	*KeyOptions     `json:"key" toml:"key"`         // Options related to key creation.
	*KeyringOptions `json:"keyring" toml:"keyring"` // Options related to keyring configuration.
	*LogOptions     `json:"log" toml:"log"`         // Options related to logging.
	*PageOptions    `json:"page" toml:"page"`       // Options related to pagination.
	*QueryOptions   `json:"query" toml:"query"`     // Options related to querying.
	*TxOptions      `json:"tx" toml:"tx"`           // Options related to transactions.
}

// New creates and returns a new instance of Options with all fields initialized to nil.
func New() *Options {
	return &Options{}
}

// NewDefault creates and returns a new instance of Options with default values for all option structs.
func NewDefault() *Options {
	return &Options{
		KeyOptions:     NewDefaultKey(),     // Initializes with default KeyOptions.
		KeyringOptions: NewDefaultKeyring(), // Initializes with default KeyringOptions.
		LogOptions:     NewDefaultLog(),     // Initializes with default LogOptions.
		PageOptions:    NewDefaultPage(),    // Initializes with default PageOptions.
		QueryOptions:   NewDefaultQuery(),   // Initializes with default QueryOptions.
		TxOptions:      NewDefaultTx(),      // Initializes with default TxOptions.
	}
}

// WithKeyOptions sets the KeyOptions for the Options and returns the updated Options.
func (o *Options) WithKeyOptions(v *KeyOptions) *Options {
	o.KeyOptions = v
	return o
}

// WithKeyringOptions sets the KeyringOptions for the Options and returns the updated Options.
func (o *Options) WithKeyringOptions(v *KeyringOptions) *Options {
	o.KeyringOptions = v
	return o
}

// WithLogOptions sets the LogOptions for the Options and returns the updated Options.
func (o *Options) WithLogOptions(v *LogOptions) *Options {
	o.LogOptions = v
	return o
}

// WithPageOptions sets the PageOptions for the Options and returns the updated Options.
func (o *Options) WithPageOptions(v *PageOptions) *Options {
	o.PageOptions = v
	return o
}

// WithQueryOptions sets the QueryOptions for the Options and returns the updated Options.
func (o *Options) WithQueryOptions(v *QueryOptions) *Options {
	o.QueryOptions = v
	return o
}

// WithTxOptions sets the TxOptions for the Options and returns the updated Options.
func (o *Options) WithTxOptions(v *TxOptions) *Options {
	o.TxOptions = v
	return o
}

// WithKeyOptionsFromCmd updates KeyOptions in the Options based on the command's flags.
func (o *Options) WithKeyOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewKeyOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKeyOptions(opts), nil
}

// WithKeyringOptionsFromCmd updates KeyringOptions in the Options based on the command's flags.
func (o *Options) WithKeyringOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewKeyringOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKeyringOptions(opts), nil
}

// WithLogOptionsFromCmd updates LogOptions in the Options based on the command's flags.
func (o *Options) WithLogOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewLogOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithLogOptions(opts), nil
}

// WithPageOptionsFromCmd updates PageOptions in the Options based on the command's flags.
func (o *Options) WithPageOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewPageOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithPageOptions(opts), nil
}

// WithQueryOptionsFromCmd updates QueryOptions in the Options based on the command's flags.
func (o *Options) WithQueryOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewQueryOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithQueryOptions(opts), nil
}

// WithTxOptionsFromCmd updates TxOptions in the Options based on the command's flags.
func (o *Options) WithTxOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewTxOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithTxOptions(opts), nil
}

// NewOptionsFromCmd creates and returns an Options instance populated with values from the command's flags.
func NewOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	// Retrieves KeyOptions from command flags.
	keyOpts, err := NewKeyOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves KeyringOptions from command flags.
	keyringOpts, err := NewKeyringOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves LogOptions from command flags.
	logOpts, err := NewLogOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves PageOptions from command flags.
	pageOpts, err := NewPageOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves QueryOptions from command flags.
	queryOpts, err := NewQueryOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves TxOptions from command flags.
	txOpts, err := NewTxOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Options instance populated with the retrieved flag values.
	return &Options{
		KeyOptions:     keyOpts,
		KeyringOptions: keyringOpts,
		LogOptions:     logOpts,
		PageOptions:    pageOpts,
		QueryOptions:   queryOpts,
		TxOptions:      txOpts,
	}, nil
}
