package options

import (
	"time"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/client/http"

	"github.com/sentinel-official/sentinel-go-sdk/v1/utils"
)

// Default values for query options.
const (
	DefaultQueryMaxRetries = 15
	DefaultQueryRPCAddr    = "https://rpc.sentinel.co:443"
	DefaultQueryTimeout    = 15 * time.Second
)

// QueryOptions represents options for making queries.
type QueryOptions struct {
	Height         int64         `json:"height,omitempty"`           // Height is the block height at which the query is to be performed.
	MaxRetries     int           `json:"max_retries,omitempty"`      // MaxRetries is the maximum number of retries for the query.
	PageCountTotal bool          `json:"page_count_total,omitempty"` // PageCountTotal indicates whether to include total count in paged queries.
	PageKey        []byte        `json:"page_key,omitempty"`         // PageKey is the key for pagination.
	PageLimit      uint64        `json:"page_limit,omitempty"`       // PageLimit is the maximum number of results per page.
	PageOffset     uint64        `json:"page_offset,omitempty"`      // PageOffset is the offset for pagination.
	PageReverse    bool          `json:"page_reverse,omitempty"`     // PageReverse indicates whether to reverse the order of results in pagination.
	Prove          bool          `json:"prove,omitempty"`            // Prove indicates whether to include proof in query results.
	RPCAddr        string        `json:"rpc_addr,omitempty"`         // RPCAddr is the address of the RPC server.
	Timeout        time.Duration `json:"timeout,omitempty"`          // Timeout is the maximum duration for the query to be executed.
}

// Query creates a new QueryOptions instance with default values.
func Query() *QueryOptions {
	return &QueryOptions{
		MaxRetries: DefaultQueryMaxRetries,
		RPCAddr:    DefaultQueryRPCAddr,
		Timeout:    DefaultQueryTimeout,
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

// WithPageCountTotal sets the PageCountTotal field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithPageCountTotal(v bool) *QueryOptions {
	q.PageCountTotal = v
	return q
}

// WithPageKey sets the PageKey field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithPageKey(v []byte) *QueryOptions {
	q.PageKey = v
	return q
}

// WithPageLimit sets the PageLimit field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithPageLimit(v uint64) *QueryOptions {
	q.PageLimit = v
	return q
}

// WithPageOffset sets the PageOffset field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithPageOffset(v uint64) *QueryOptions {
	q.PageOffset = v
	return q
}

// WithPageReverse sets the PageReverse field and returns the modified QueryOptions instance.
func (q *QueryOptions) WithPageReverse(v bool) *QueryOptions {
	q.PageReverse = v
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

// PageRequest creates a new PageRequest with the configured options.
func (q *QueryOptions) PageRequest() *query.PageRequest {
	return &query.PageRequest{
		Key:        q.PageKey,
		Offset:     q.PageOffset,
		Limit:      q.PageLimit,
		CountTotal: q.PageCountTotal,
		Reverse:    q.PageReverse,
	}
}
