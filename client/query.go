package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/bytes"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/v2fly/v2ray-core/v5/common/retry"
)

// ABCIQueryWithOptions performs an ABCI query with configurable options.
func (c *Client) ABCIQueryWithOptions(ctx context.Context, path string, data bytes.HexBytes, opts *Options) (*abcitypes.ResponseQuery, error) {
	var result *coretypes.ResultABCIQuery

	fn := func() error {
		// Get the RPC client from the provided options.
		rpcClient, err := opts.Client()
		if err != nil {
			return err
		}

		// Perform the ABCI query with the given options.
		result, err = rpcClient.ABCIQueryWithOptions(ctx, path, data, opts.ABCIQueryOptions())
		if err != nil {
			return err
		}

		// If query is successful, return nil to continue.
		return nil
	}

	// Convert retry delay from time.Duration to milliseconds.
	retryDelay := uint32(opts.GetRetryDelay() / time.Millisecond)

	// Retry the query based on the maximum retry attempts and delay specified in options.
	if err := retry.Timed(opts.GetMaxRetries(), retryDelay).On(fn); err != nil {
		return nil, err
	}

	// Return nil if no result was produced.
	if result == nil {
		return nil, nil
	}

	// Return the final response from the query.
	return &result.Response, nil
}

// QueryKey performs an ABCI query for a specific key in a store.
func (c *Client) QueryKey(ctx context.Context, store string, data bytes.HexBytes, opts *Options) (*abcitypes.ResponseQuery, error) {
	// Construct the path for querying a key in the store.
	path := fmt.Sprintf("/store/%s/key", store)

	// Delegate the ABCI query to ABCIQueryWithOptions.
	return c.ABCIQueryWithOptions(ctx, path, data, opts)
}

// QuerySubspace performs an ABCI query for a subspace in a store.
func (c *Client) QuerySubspace(ctx context.Context, store string, data bytes.HexBytes, opts *Options) (*abcitypes.ResponseQuery, error) {
	// Construct the path for querying a subspace in the store.
	path := fmt.Sprintf("/store/%s/subspace", store)

	// Delegate the ABCI query to ABCIQueryWithOptions.
	return c.ABCIQueryWithOptions(ctx, path, data, opts)
}

// QueryGRPC performs a gRPC query using ABCI with configurable options.
// It marshals the request, queries with ABCI, and unmarshals the response.
func (c *Client) QueryGRPC(ctx context.Context, method string, req, resp codec.ProtoMarshaler, opts *Options) error {
	// Marshal the gRPC request.
	data, err := c.Marshal(req)
	if err != nil {
		return err
	}

	// Perform ABCI query with options.
	reply, err := c.ABCIQueryWithOptions(ctx, method, data, opts)
	if err != nil {
		return err
	}

	// Check for a nil reply.
	if reply == nil {
		return errors.New("nil reply")
	}

	// Unmarshal the ABCI response value into the provided response object.
	if err := c.Unmarshal(reply.Value, resp); err != nil {
		return err
	}

	// Return nil on success.
	return nil
}
