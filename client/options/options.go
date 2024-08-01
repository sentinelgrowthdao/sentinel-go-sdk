package options

import (
	"github.com/spf13/cobra"
)

// Options aggregates all the individual option structs.
type Options struct {
	*KeyOptions     // Options related to key creation.
	*KeyringOptions // Options related to keyring configuration.
	*PageOptions    // Options related to pagination.
	*QueryOptions   // Options related to querying.
	*TxOptions      // Options related to transactions.
}

// NewOptions creates and returns a new instance of Options with all fields initialized to nil.
func NewOptions() *Options {
	return &Options{}
}

// NewDefaultOptions creates and returns a new instance of Options with default values.
func NewDefaultOptions() *Options {
	return &Options{
		KeyOptions:     NewDefaultKeyOptions(),     // Initializes with default KeyOptions.
		KeyringOptions: NewDefaultKeyringOptions(), // Initializes with default KeyringOptions.
		PageOptions:    NewDefaultPageOptions(),    // Initializes with default PageOptions.
		QueryOptions:   NewDefaultQueryOptions(),   // Initializes with default QueryOptions.
		TxOptions:      NewDefaultTxOptions(),      // Initializes with default TxOptions.
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

// WithKeyOptionsFromCmd updates KeyOptions in the Options from the given command's flags.
func (o *Options) WithKeyOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewKeyOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKeyOptions(opts), nil
}

// WithKeyringOptionsFromCmd updates KeyringOptions in the Options from the given command's flags.
func (o *Options) WithKeyringOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewKeyringOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKeyringOptions(opts), nil
}

// WithPageOptionsFromCmd updates PageOptions in the Options from the given command's flags.
func (o *Options) WithPageOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewPageOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithPageOptions(opts), nil
}

// WithQueryOptionsFromCmd updates QueryOptions in the Options from the given command's flags.
func (o *Options) WithQueryOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewQueryOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithQueryOptions(opts), nil
}

// WithTxOptionsFromCmd updates TxOptions in the Options from the given command's flags.
func (o *Options) WithTxOptionsFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewTxOptionsFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithTxOptions(opts), nil
}

// NewOptionsFromCmd creates and returns an Options instance populated with values from the given command's flags.
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
		PageOptions:    pageOpts,
		QueryOptions:   queryOpts,
		TxOptions:      txOpts,
	}, nil
}
