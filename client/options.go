package client

import (
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/options"
)

// Options aggregates all the individual option structs for a comprehensive configuration.
type Options struct {
	*options.Key     `json:"key" toml:"key"`         // Options related to key creation.
	*options.Keyring `json:"keyring" toml:"keyring"` // Options related to keyring configuration.
	*options.Page    `json:"page" toml:"page"`       // Options related to pagination.
	*options.Query   `json:"query" toml:"query"`     // Options related to querying.
	*options.Tx      `json:"tx" toml:"tx"`           // Options related to transactions.
}

// NewOptions creates and returns a new instance of Options.
func NewOptions() *Options {
	return &Options{}
}

// WithKey sets the Key for the Options and returns the updated Options.
func (o *Options) WithKey(v *options.Key) *Options {
	o.Key = v
	return o
}

// WithKeyring sets the Keyring for the Options and returns the updated Options.
func (o *Options) WithKeyring(v *options.Keyring) *Options {
	o.Keyring = v
	return o
}

// WithPage sets the Page for the Options and returns the updated Options.
func (o *Options) WithPage(v *options.Page) *Options {
	o.Page = v
	return o
}

// WithQuery sets the Query for the Options and returns the updated Options.
func (o *Options) WithQuery(v *options.Query) *Options {
	o.Query = v
	return o
}

// WithTx sets the Tx for the Options and returns the updated Options.
func (o *Options) WithTx(v *options.Tx) *Options {
	o.Tx = v
	return o
}

// WithKeyFromCmd updates Key in the Options based on the command's flags.
func (o *Options) WithKeyFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := options.NewKeyFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKey(opts), nil
}

// WithKeyringFromCmd updates Keyring in the Options based on the command's flags.
func (o *Options) WithKeyringFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := options.NewKeyringFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKeyring(opts), nil
}

// WithPageFromCmd updates Page in the Options based on the command's flags.
func (o *Options) WithPageFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := options.NewPageFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithPage(opts), nil
}

// WithQueryFromCmd updates Query in the Options based on the command's flags.
func (o *Options) WithQueryFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := options.NewQueryFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithQuery(opts), nil
}

// WithTxFromCmd updates Tx in the Options based on the command's flags.
func (o *Options) WithTxFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := options.NewTxFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithTx(opts), nil
}

// Validate ensures all option fields are valid.
func (o *Options) Validate() error {
	if o.Key != nil {
		if err := o.Key.Validate(); err != nil {
			return err
		}
	}
	if o.Keyring != nil {
		if err := o.Keyring.Validate(); err != nil {
			return err
		}
	}
	if o.Page != nil {
		if err := o.Page.Validate(); err != nil {
			return err
		}
	}
	if o.Query != nil {
		if err := o.Query.Validate(); err != nil {
			return err
		}
	}
	if o.Tx != nil {
		if err := o.Tx.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// NewFromCmd creates and returns an Options instance populated with values from the command's flags.
func NewFromCmd(cmd *cobra.Command) (*Options, error) {
	// Retrieves Key from command flags.
	key, err := options.NewKeyFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Keyring from command flags.
	keyring, err := options.NewKeyringFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Page from command flags.
	page, err := options.NewPageFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Query from command flags.
	query, err := options.NewQueryFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Tx from command flags.
	tx, err := options.NewTxFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Options instance populated with the retrieved flag values.
	return &Options{
		Key:     key,
		Keyring: keyring,
		Page:    page,
		Query:   query,
		Tx:      tx,
	}, nil
}
