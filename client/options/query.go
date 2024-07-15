package options

import (
	"time"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/spf13/cobra"

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

// AddQueryFlagsToCmd adds query-related flags to the given cobra command.
func AddQueryFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().Int64("query.height", 0, "Block height for the query.")
	cmd.Flags().Int("query.max-retries", DefaultQueryMaxRetries, "Maximum number of retries for the query.")
	cmd.Flags().Bool("query.page-count-total", false, "Include total count in paged queries.")
	cmd.Flags().BytesBase64("query.page-key", nil, "Base64-encoded key for pagination.")
	cmd.Flags().Uint64("query.page-limit", 0, "Maximum number of results per page.")
	cmd.Flags().Uint64("query.page-offset", 0, "Offset for pagination.")
	cmd.Flags().Bool("query.page-reverse", false, "Reverse the order of results in pagination.")
	cmd.Flags().Bool("query.prove", false, "Include proof in query results.")
	cmd.Flags().String("query.rpc-addr", DefaultQueryRPCAddr, "Address of the RPC server.")
	cmd.Flags().Duration("query.timeout", DefaultQueryTimeout, "Maximum duration for the query execution.")
}

// NewQueryOptionsFromCmd creates and returns QueryOptions from the given cobra command's flags.
func NewQueryOptionsFromCmd(cmd *cobra.Command) (*QueryOptions, error) {
	// Retrieve the value of the "query.height" flag.
	height, err := cmd.Flags().GetInt64("query.height")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.max-retries" flag.
	maxRetries, err := cmd.Flags().GetInt("query.max-retries")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.page-count-total" flag.
	pageCountTotal, err := cmd.Flags().GetBool("query.page-count-total")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.page-key" flag.
	pageKey, err := cmd.Flags().GetBytesBase64("query.page-key")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.page-limit" flag.
	pageLimit, err := cmd.Flags().GetUint64("query.page-limit")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.page-offset" flag.
	pageOffset, err := cmd.Flags().GetUint64("query.page-offset")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.page-reverse" flag.
	pageReverse, err := cmd.Flags().GetBool("query.page-reverse")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.prove" flag.
	prove, err := cmd.Flags().GetBool("query.prove")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.rpc-addr" flag.
	rpcAddr, err := cmd.Flags().GetString("query.rpc-addr")
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.timeout" flag.
	timeout, err := cmd.Flags().GetDuration("query.timeout")
	if err != nil {
		return nil, err
	}

	// Return a new QueryOptions instance populated with the retrieved flag values.
	return &QueryOptions{
		Height:         height,
		MaxRetries:     maxRetries,
		PageCountTotal: pageCountTotal,
		PageKey:        pageKey,
		PageLimit:      pageLimit,
		PageOffset:     pageOffset,
		PageReverse:    pageReverse,
		Prove:          prove,
		RPCAddr:        rpcAddr,
		Timeout:        timeout,
	}, nil
}
