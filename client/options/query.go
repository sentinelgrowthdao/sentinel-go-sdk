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

// Page represents page-related options.
type Page struct {
	CountTotal bool   `json:"count_total" toml:"count_total"` // CountTotal indicates whether to include total count in paged queries.
	Key        []byte `json:"key" toml:"key"`                 // Key is the key for page.
	Limit      uint64 `json:"limit" toml:"limit"`             // Limit is the maximum number of results per page.
	Offset     uint64 `json:"offset" toml:"offset"`           // Offset is the offset for page.
	Reverse    bool   `json:"reverse" toml:"reverse"`         // Reverse indicates whether to reverse the order of results in page.
}

// NewPage creates a new Page instance with default values.
func NewPage() *Page {
	return &Page{
		Limit: flags.DefaultPageLimit,
	}
}

// WithCountTotal sets the CountTotal field and returns the modified Page instance.
func (p *Page) WithCountTotal(v bool) *Page {
	p.CountTotal = v
	return p
}

// WithKey sets the Key field and returns the modified Page instance.
func (p *Page) WithKey(v []byte) *Page {
	p.Key = v
	return p
}

// WithLimit sets the Limit field and returns the modified Page instance.
func (p *Page) WithLimit(v uint64) *Page {
	p.Limit = v
	return p
}

// WithOffset sets the Offset field and returns the modified Page instance.
func (p *Page) WithOffset(v uint64) *Page {
	p.Offset = v
	return p
}

// WithReverse sets the Reverse field and returns the modified Page instance.
func (p *Page) WithReverse(v bool) *Page {
	p.Reverse = v
	return p
}

// PageRequest creates a new PageRequest with the configured options.
func (p *Page) PageRequest() *query.PageRequest {
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

// NewPageFromCmd creates and returns Page from the given cobra command's flags.
func NewPageFromCmd(cmd *cobra.Command) (*Page, error) {
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

	// Return a new Page instance populated with the retrieved flag values.
	return &Page{
		CountTotal: countTotal,
		Key:        key,
		Limit:      limit,
		Offset:     offset,
		Reverse:    reverse,
	}, nil
}

// Query represents options for making queries.
type Query struct {
	Height     int64         `json:"height" toml:"height"`           // Height is the block height at which the query is to be performed.
	MaxRetries int           `json:"max_retries" toml:"max_retries"` // MaxRetries is the maximum number of retries for the query.
	Prove      bool          `json:"prove" toml:"prove"`             // Prove indicates whether to include proof in query results.
	RPCAddr    string        `json:"rpc_addr" toml:"rpc_addr"`       // RPCAddr is the address of the RPC server.
	Timeout    time.Duration `json:"timeout" toml:"timeout"`         // Timeout is the maximum duration for the query to be executed.
}

// NewQuery creates a new Query instance with default values.
func NewQuery() *Query {
	return &Query{
		MaxRetries: flags.DefaultQueryMaxRetries,
		RPCAddr:    flags.DefaultQueryRPCAddr,
		Timeout:    flags.DefaultQueryTimeout,
	}
}

// WithHeight sets the Height field and returns the modified Query instance.
func (q *Query) WithHeight(v int64) *Query {
	q.Height = v
	return q
}

// WithMaxRetries sets the MaxRetries field and returns the modified Query instance.
func (q *Query) WithMaxRetries(v int) *Query {
	q.MaxRetries = v
	return q
}

// WithProve sets the Prove field and returns the modified Query instance.
func (q *Query) WithProve(v bool) *Query {
	q.Prove = v
	return q
}

// WithRPCAddr sets the RPCAddr field and returns the modified Query instance.
func (q *Query) WithRPCAddr(v string) *Query {
	q.RPCAddr = v
	return q
}

// WithTimeout sets the Timeout field and returns the modified Query instance.
func (q *Query) WithTimeout(v time.Duration) *Query {
	q.Timeout = v
	return q
}

// ABCIQueryOptions converts Query to ABCIQueryOptions.
func (q *Query) ABCIQueryOptions() client.ABCIQueryOptions {
	return client.ABCIQueryOptions{
		Height: q.Height,
		Prove:  q.Prove,
	}
}

// Client creates a new HTTP client with the configured options.
func (q *Query) Client() (*http.HTTP, error) {
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

// NewQueryFromCmd creates and returns Query from the given cobra command's flags.
func NewQueryFromCmd(cmd *cobra.Command) (*Query, error) {
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

	// Return a new Query instance populated with the retrieved flag values.
	return &Query{
		Height:     height,
		MaxRetries: maxRetries,
		Prove:      prove,
		RPCAddr:    rpcAddr,
		Timeout:    timeout,
	}, nil
}
