package options

import (
	"time"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/cmd/flags"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
)

// PageOptions represents page-related options.
type PageOptions struct {
	CountTotal bool   `json:"count_total" toml:"count_total"` // CountTotal indicates whether to include total count in paged queries.
	Key        []byte `json:"key" toml:"key"`                 // Key is the key for page.
	Limit      uint64 `json:"limit" toml:"limit"`             // Limit is the maximum number of results per page.
	Offset     uint64 `json:"offset" toml:"offset"`           // Offset is the offset for page.
	Reverse    bool   `json:"reverse" toml:"reverse"`         // Reverse indicates whether to reverse the order of results in page.
}

// NewDefaultPage creates a new PageOptions instance with default values.
func NewDefaultPage() *PageOptions {
	return &PageOptions{
		Limit: flags.DefaultPageLimit,
	}
}

// WithCountTotal sets the CountTotal field and returns the modified PageOptions instance.
func (p *PageOptions) WithCountTotal(v bool) *PageOptions {
	p.CountTotal = v
	return p
}

// WithKey sets the Key field and returns the modified PageOptions instance.
func (p *PageOptions) WithKey(v []byte) *PageOptions {
	p.Key = v
	return p
}

// WithLimit sets the Limit field and returns the modified PageOptions instance.
func (p *PageOptions) WithLimit(v uint64) *PageOptions {
	p.Limit = v
	return p
}

// WithOffset sets the Offset field and returns the modified PageOptions instance.
func (p *PageOptions) WithOffset(v uint64) *PageOptions {
	p.Offset = v
	return p
}

// WithReverse sets the Reverse field and returns the modified PageOptions instance.
func (p *PageOptions) WithReverse(v bool) *PageOptions {
	p.Reverse = v
	return p
}

// PageRequest creates a new PageRequest with the configured options.
func (p *PageOptions) PageRequest() *query.PageRequest {
	return &query.PageRequest{
		Key:        p.Key,
		Offset:     p.Offset,
		Limit:      p.Limit,
		CountTotal: p.CountTotal,
		Reverse:    p.Reverse,
	}
}

// AddPageFlagsToCmd adds page-related flags to the given cobra command.
func AddPageFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagPageCountTotal(cmd)
	flags.SetFlagPageKey(cmd)
	flags.SetFlagPageLimit(cmd)
	flags.SetFlagPageOffset(cmd)
	flags.SetFlagPageReverse(cmd)
}

// NewPageOptionsFromCmd creates and returns PageOptions from the given cobra command's flags.
func NewPageOptionsFromCmd(cmd *cobra.Command) (*PageOptions, error) {
	// Retrieve the value of the "page.count-total" flag.
	countTotal, err := flags.GetPageCountTotalFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.key" flag.
	key, err := flags.GetPageKeyFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.limit" flag.
	limit, err := flags.GetPageLimitFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.offset" flag.
	offset, err := flags.GetPageOffsetFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.reverse" flag.
	reverse, err := flags.GetPageReverseFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new PageOptions instance populated with the retrieved flag values.
	return &PageOptions{
		CountTotal: countTotal,
		Key:        key,
		Limit:      limit,
		Offset:     offset,
		Reverse:    reverse,
	}, nil
}

// QueryOptions represents options for making queries.
type QueryOptions struct {
	Height     int64         `json:"height" toml:"height"`           // Height is the block height at which the query is to be performed.
	MaxRetries int           `json:"max_retries" toml:"max_retries"` // MaxRetries is the maximum number of retries for the query.
	Prove      bool          `json:"prove" toml:"prove"`             // Prove indicates whether to include proof in query results.
	RPCAddr    string        `json:"rpc_addr" toml:"rpc_addr"`       // RPCAddr is the address of the RPC server.
	Timeout    time.Duration `json:"timeout" toml:"timeout"`         // Timeout is the maximum duration for the query to be executed.
}

// NewDefaultQuery creates a new QueryOptions instance with default values.
func NewDefaultQuery() *QueryOptions {
	return &QueryOptions{
		MaxRetries: flags.DefaultQueryMaxRetries,
		RPCAddr:    flags.DefaultQueryRPCAddr,
		Timeout:    flags.DefaultQueryTimeout,
	}
}

// WithHeight sets the Height field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithHeight(v int64) *QueryOptions {
	q.Height = v
	return q
}

// WithMaxRetries sets the MaxRetries field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithMaxRetries(v int) *QueryOptions {
	q.MaxRetries = v
	return q
}

// WithProve sets the Prove field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithProve(v bool) *QueryOptions {
	q.Prove = v
	return q
}

// WithRPCAddr sets the RPCAddr field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithRPCAddr(v string) *QueryOptions {
	q.RPCAddr = v
	return q
}

// WithTimeout sets the Timeout field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithTimeout(v time.Duration) *QueryOptions {
	q.Timeout = v
	return q
}

// ABCIQueryOptions converts QueryOptions to ABCIQueryOptions.
func (q *QueryOptions) ABCIQueryOptions() client.ABCIQueryOptions {
	return client.ABCIQueryOptions{
		Height: q.Height,
		Prove:  q.Prove,
	}
}

// Client creates a new HTTP client with the configured options.
func (q *QueryOptions) Client() (*http.HTTP, error) {
	return http.NewWithTimeout(q.RPCAddr, "/websocket", utils.UIntSecondsFromDuration(q.Timeout))
}

// AddQueryFlagsToCmd adds query-related flags to the given cobra command.
func AddQueryFlagsToCmd(cmd *cobra.Command) {
	flags.SetFlagQueryHeight(cmd)
	flags.SetFlagQueryMaxRetries(cmd)
	flags.SetFlagQueryProve(cmd)
	flags.SetFlagQueryRPCAddr(cmd)
	flags.SetFlagQueryTimeout(cmd)
}

// NewQueryOptionsFromCmd creates and returns QueryOptions from the given cobra command's flags.
func NewQueryOptionsFromCmd(cmd *cobra.Command) (*QueryOptions, error) {
	// Retrieve the height flag value from the command.
	height, err := flags.GetQueryHeightFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the max retries flag value from the command.
	maxRetries, err := flags.GetQueryMaxRetriesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the prove flag value from the command.
	prove, err := flags.GetQueryProveFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the RPC address flag value from the command.
	rpcAddr, err := flags.GetQueryRPCAddrFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the timeout flag value from the command.
	timeout, err := flags.GetQueryTimeoutFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Return a new QueryOptions instance populated with the retrieved flag values.
	return &QueryOptions{
		Height:     height,
		MaxRetries: maxRetries,
		Prove:      prove,
		RPCAddr:    rpcAddr,
		Timeout:    timeout,
	}, nil
}
