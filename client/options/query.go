package options

import (
	"time"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/utils"
)

// Default values for page and query options.
const (
	DefaultPageLimit = 25

	DefaultQueryMaxRetries = 15
	DefaultQueryRPCAddr    = "https://rpc.sentinel.co:443"
	DefaultQueryTimeout    = 15 * time.Second
)

// PageOptions represents page-related options.
type PageOptions struct {
	CountTotal bool   `json:"count_total,omitempty"` // CountTotal indicates whether to include total count in paged queries.
	Key        []byte `json:"key,omitempty"`         // Key is the key for page.
	Limit      uint64 `json:"limit,omitempty"`       // Limit is the maximum number of results per page.
	Offset     uint64 `json:"offset,omitempty"`      // Offset is the offset for page.
	Reverse    bool   `json:"reverse,omitempty"`     // Reverse indicates whether to reverse the order of results in page.
}

// NewDefaultPageOptions creates a new PageOptions instance with default values.
func NewDefaultPageOptions() *PageOptions {
	return &PageOptions{
		Limit: DefaultPageLimit,
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

// GetPageCountFromCmd retrieves the "page.count-total" flag value from the command.
func GetPageCountFromCmd(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("page.count-total")
}

// GetPageKeyFromCmd retrieves the "page.key" flag value from the command.
func GetPageKeyFromCmd(cmd *cobra.Command) ([]byte, error) {
	return cmd.Flags().GetBytesBase64("page.key")
}

// GetPageLimitFromCmd retrieves the "page.limit" flag value from the command.
func GetPageLimitFromCmd(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("page.limit")
}

// GetPageOffsetFromCmd retrieves the "page.offset" flag value from the command.
func GetPageOffsetFromCmd(cmd *cobra.Command) (uint64, error) {
	return cmd.Flags().GetUint64("page.offset")
}

// GetPageReverseFromCmd retrieves the "page.reverse" flag value from the command.
func GetPageReverseFromCmd(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("page.reverse")
}

// SetFlagPageCount adds the "page.count-total" flag to the command.
func SetFlagPageCount(cmd *cobra.Command) {
	cmd.Flags().Bool("page.count-total", false, "Include total count in paged queries.")
}

// SetFlagPageKey adds the "page.key" flag to the command.
func SetFlagPageKey(cmd *cobra.Command) {
	cmd.Flags().BytesBase64("page.key", nil, "Base64-encoded key for page.")
}

// SetFlagPageLimit adds the "page.limit" flag to the command.
func SetFlagPageLimit(cmd *cobra.Command) {
	cmd.Flags().Uint64("page.limit", DefaultPageLimit, "Maximum number of results per page.")
}

// SetFlagPageOffset adds the "page.offset" flag to the command.
func SetFlagPageOffset(cmd *cobra.Command) {
	cmd.Flags().Uint64("page.offset", 0, "Offset for page.")
}

// SetFlagPageReverse adds the "page.reverse" flag to the command.
func SetFlagPageReverse(cmd *cobra.Command) {
	cmd.Flags().Bool("page.reverse", false, "Reverse the order of results in page.")
}

// AddPageFlagsToCmd adds page-related flags to the given cobra command.
func AddPageFlagsToCmd(cmd *cobra.Command) {
	SetFlagPageCount(cmd)
	SetFlagPageKey(cmd)
	SetFlagPageLimit(cmd)
	SetFlagPageOffset(cmd)
	SetFlagPageReverse(cmd)
}

// NewPageOptionsFromCmd creates and returns PageOptions from the given cobra command's flags.
func NewPageOptionsFromCmd(cmd *cobra.Command) (*PageOptions, error) {
	// Retrieve the value of the "page.count-total" flag.
	countTotal, err := GetPageCountFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.key" flag.
	key, err := GetPageKeyFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.limit" flag.
	limit, err := GetPageLimitFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.offset" flag.
	offset, err := GetPageOffsetFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "page.reverse" flag.
	reverse, err := GetPageReverseFromCmd(cmd)
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
	Height     int64         `json:"height,omitempty"`      // Height is the block height at which the query is to be performed.
	MaxRetries int           `json:"max_retries,omitempty"` // MaxRetries is the maximum number of retries for the query.
	Prove      bool          `json:"prove,omitempty"`       // Prove indicates whether to include proof in query results.
	RPCAddr    string        `json:"rpc_addr,omitempty"`    // RPCAddr is the address of the RPC server.
	Timeout    time.Duration `json:"timeout,omitempty"`     // Timeout is the maximum duration for the query to be executed.
}

// NewDefaultQueryOptions creates a new QueryOptions instance with default values.
func NewDefaultQueryOptions() *QueryOptions {
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

// GetHeightFromCmd retrieves the "query.height" flag value from the command.
func GetHeightFromCmd(cmd *cobra.Command) (int64, error) {
	return cmd.Flags().GetInt64("query.height")
}

// GetMaxRetriesFromCmd retrieves the "query.max-retries" flag value from the command.
func GetMaxRetriesFromCmd(cmd *cobra.Command) (int, error) {
	return cmd.Flags().GetInt("query.max-retries")
}

// GetProveFromCmd retrieves the "query.prove" flag value from the command.
func GetProveFromCmd(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool("query.prove")
}

// GetRPCAddrFromCmd retrieves the "query.rpc-addr" flag value from the command.
func GetRPCAddrFromCmd(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("query.rpc-addr")
}

// GetTimeoutFromCmd retrieves the "query.timeout" flag value from the command.
func GetTimeoutFromCmd(cmd *cobra.Command) (time.Duration, error) {
	return cmd.Flags().GetDuration("query.timeout")
}

// SetFlagHeight adds the "query.height" flag to the command.
func SetFlagHeight(cmd *cobra.Command) {
	cmd.Flags().Int64("query.height", 0, "Block height at which the query is to be performed.")
}

// SetFlagMaxRetries adds the "query.max-retries" flag to the command.
func SetFlagMaxRetries(cmd *cobra.Command) {
	cmd.Flags().Int("query.max-retries", DefaultQueryMaxRetries, "Maximum number of retries for the query.")
}

// SetFlagProve adds the "query.prove" flag to the command.
func SetFlagProve(cmd *cobra.Command) {
	cmd.Flags().Bool("query.prove", false, "Include proof in query results.")
}

// SetFlagRPCAddr adds the "query.rpc-addr" flag to the command.
func SetFlagRPCAddr(cmd *cobra.Command) {
	cmd.Flags().String("query.rpc-addr", DefaultQueryRPCAddr, "Address of the RPC server.")
}

// SetFlagTimeout adds the "query.timeout" flag to the command.
func SetFlagTimeout(cmd *cobra.Command) {
	cmd.Flags().Duration("query.timeout", DefaultQueryTimeout, "Maximum duration for the query to be executed.")
}

// AddQueryFlagsToCmd adds query-related flags to the given cobra command.
func AddQueryFlagsToCmd(cmd *cobra.Command) {
	SetFlagHeight(cmd)
	SetFlagMaxRetries(cmd)
	SetFlagProve(cmd)
	SetFlagRPCAddr(cmd)
	SetFlagTimeout(cmd)
}

// NewQueryOptionsFromCmd creates and returns QueryOptions from the given cobra command's flags.
func NewQueryOptionsFromCmd(cmd *cobra.Command) (*QueryOptions, error) {
	// Retrieve the value of the "query.height" flag.
	height, err := GetHeightFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.max-retries" flag.
	maxRetries, err := GetMaxRetriesFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.prove" flag.
	prove, err := GetProveFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.rpc-addr" flag.
	rpcAddr, err := GetRPCAddrFromCmd(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the value of the "query.timeout" flag.
	timeout, err := GetTimeoutFromCmd(cmd)
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
