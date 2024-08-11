package options

import (
	"github.com/spf13/cobra"
)

// Options aggregates all the individual option structs for a comprehensive configuration.
type Options struct {
	*Key     `json:"key" toml:"key"`         // Options related to key creation.
	*Keyring `json:"keyring" toml:"keyring"` // Options related to keyring configuration.
	*Log     `json:"log" toml:"log"`         // Options related to logging.
	*Page    `json:"page" toml:"page"`       // Options related to pagination.
	*Query   `json:"query" toml:"query"`     // Options related to querying.
	*Tx      `json:"tx" toml:"tx"`           // Options related to transactions.
}

// New creates and returns a new instance of Options with default values for all option structs.
func New() *Options {
	return &Options{
		Key:     NewKey(),     // Initializes with default Key.
		Keyring: NewKeyring(), // Initializes with default Keyring.
		Log:     NewLog(),     // Initializes with default Log.
		Page:    NewPage(),    // Initializes with default Page.
		Query:   NewQuery(),   // Initializes with default Query.
		Tx:      NewTx(),      // Initializes with default Tx.
	}
}

// WithKey sets the Key for the Options and returns the updated Options.
func (o *Options) WithKey(v *Key) *Options {
	o.Key = v
	return o
}

// WithKeyring sets the Keyring for the Options and returns the updated Options.
func (o *Options) WithKeyring(v *Keyring) *Options {
	o.Keyring = v
	return o
}

// WithLog sets the Log for the Options and returns the updated Options.
func (o *Options) WithLog(v *Log) *Options {
	o.Log = v
	return o
}

// WithPage sets the Page for the Options and returns the updated Options.
func (o *Options) WithPage(v *Page) *Options {
	o.Page = v
	return o
}

// WithQuery sets the Query for the Options and returns the updated Options.
func (o *Options) WithQuery(v *Query) *Options {
	o.Query = v
	return o
}

// WithTx sets the Tx for the Options and returns the updated Options.
func (o *Options) WithTx(v *Tx) *Options {
	o.Tx = v
	return o
}

// WithKeyFromCmd updates Key in the Options based on the command's flags.
func (o *Options) WithKeyFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewKeyFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKey(opts), nil
}

// WithKeyringFromCmd updates Keyring in the Options based on the command's flags.
func (o *Options) WithKeyringFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewKeyringFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithKeyring(opts), nil
}

// WithLogFromCmd updates Log in the Options based on the command's flags.
func (o *Options) WithLogFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewLogFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithLog(opts), nil
}

// WithPageFromCmd updates Page in the Options based on the command's flags.
func (o *Options) WithPageFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewPageFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithPage(opts), nil
}

// WithQueryFromCmd updates Query in the Options based on the command's flags.
func (o *Options) WithQueryFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewQueryFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithQuery(opts), nil
}

// WithTxFromCmd updates Tx in the Options based on the command's flags.
func (o *Options) WithTxFromCmd(cmd *cobra.Command) (*Options, error) {
	opts, err := NewTxFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	return o.WithTx(opts), nil
}

// NewFromCmd creates and returns an Options instance populated with values from the command's flags.
func NewFromCmd(cmd *cobra.Command) (*Options, error) {
	// Retrieves Key from command flags.
	key, err := NewKeyFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Keyring from command flags.
	keyring, err := NewKeyringFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Log from command flags.
	log, err := NewLogFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Page from command flags.
	page, err := NewPageFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Query from command flags.
	query, err := NewQueryFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieves Tx from command flags.
	tx, err := NewTxFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new Options instance populated with the retrieved flag values.
	return &Options{
		Key:     key,
		Keyring: keyring,
		Log:     log,
		Page:    page,
		Query:   query,
		Tx:      tx,
	}, nil
}
