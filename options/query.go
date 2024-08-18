package options

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/sentinel-go-sdk/flags"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
)

// Query represents options for making queries.
type Query struct {
	Height     int64  `json:"height" toml:"height"`           // Height is the block height at which the query is to be performed.
	MaxRetries int    `json:"max_retries" toml:"max_retries"` // MaxRetries is the maximum number of retries for the query.
	Prove      bool   `json:"prove" toml:"prove"`             // Prove indicates whether to include proof in query results.
	RPCAddr    string `json:"rpc_addr" toml:"rpc_addr"`       // RPCAddr is the address of the RPC server.
	Timeout    string `json:"timeout" toml:"timeout"`         // Timeout is the maximum duration for the query to be executed.
}

// NewQuery creates a new Query instance with default values.
func NewQuery() *Query {
	return &Query{
		Height:     flags.DefaultQueryHeight,
		MaxRetries: flags.DefaultQueryMaxRetries,
		Prove:      flags.DefaultQueryProve,
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
	q.Timeout = v.String()
	return q
}

// GetHeight returns the block height for the query.
func (q *Query) GetHeight() int64 {
	return q.Height
}

// GetMaxRetries returns the maximum number of retries for the query.
func (q *Query) GetMaxRetries() int {
	return q.MaxRetries
}

// GetProve returns whether to include proof in query results.
func (q *Query) GetProve() bool {
	return q.Prove
}

// GetRPCAddr returns the address of the RPC server.
func (q *Query) GetRPCAddr() string {
	return q.RPCAddr
}

// GetTimeout returns the maximum duration for the query.
func (q *Query) GetTimeout() time.Duration {
	v, err := time.ParseDuration(q.Timeout)
	if err != nil {
		panic(err)
	}

	return v
}

// ValidateQueryHeight validates the Height field.
func ValidateQueryHeight(v int64) error {
	if v < 0 {
		return errors.New("height must be non-negative")
	}

	return nil
}

// ValidateQueryMaxRetries validates the MaxRetries field.
func ValidateQueryMaxRetries(v int) error {
	if v < 0 {
		return errors.New("max_retries must be non-negative")
	}

	return nil
}

// ValidateQueryRPCAddr validates the RPCAddr field.
func ValidateQueryRPCAddr(v string) error {
	if v == "" {
		return errors.New("rpc_addr must not be empty")
	}

	// Parse the URL
	addr, err := url.Parse(v)
	if err != nil {
		return errors.New("rpc_addr must be a valid URL")
	}

	// Check if the URL scheme is set
	if addr.Scheme == "" {
		return errors.New("rpc_addr must have a valid scheme (e.g., http, https)")
	}

	// Check if the URL host is set and contains a port
	if addr.Host == "" || !strings.Contains(addr.Host, ":") {
		return errors.New("rpc_addr must be a valid URL with a port")
	}

	return nil
}

// ValidateQueryTimeout validates the Timeout field.
func ValidateQueryTimeout(timeout string) error {
	if _, err := time.ParseDuration(timeout); err != nil {
		return errors.New("timeout must be a valid duration")
	}

	return nil
}

// Validate validates all the fields of the Query struct.
func (q *Query) Validate() error {
	if err := ValidateQueryHeight(q.Height); err != nil {
		return err
	}
	if err := ValidateQueryMaxRetries(q.MaxRetries); err != nil {
		return err
	}
	if err := ValidateQueryRPCAddr(q.RPCAddr); err != nil {
		return err
	}
	if err := ValidateQueryTimeout(q.Timeout); err != nil {
		return err
	}

	return nil
}

// ABCIQueryOptions converts Query to ABCIQueryOptions.
func (q *Query) ABCIQueryOptions() client.ABCIQueryOptions {
	return client.ABCIQueryOptions{
		Height: q.GetHeight(),
		Prove:  q.GetProve(),
	}
}

// Client creates a new HTTP client with the configured options.
func (q *Query) Client() (*http.HTTP, error) {
	timeout := utils.UIntSecondsFromDuration(q.GetTimeout())
	return http.NewWithTimeout(q.GetRPCAddr(), "/websocket", timeout)
}

// NewQueryFromCmd creates and returns Query from the given cobra command's flags.
func NewQueryFromCmd(cmd *cobra.Command) (*Query, error) {
	// Retrieve the height flag value from the command.
	height, err := flags.GetQueryHeight(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the max retries flag value from the command.
	maxRetries, err := flags.GetQueryMaxRetries(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the prove flag value from the command.
	prove, err := flags.GetQueryProve(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the RPC address flag value from the command.
	rpcAddr, err := flags.GetQueryRPCAddr(cmd)
	if err != nil {
		return nil, err
	}

	// Retrieve the timeout flag value from the command.
	timeout, err := flags.GetQueryTimeout(cmd)
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
