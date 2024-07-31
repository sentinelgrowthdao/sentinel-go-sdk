package client

import (
	"context"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"

	"github.com/sentinel-official/sentinel-go-sdk/v1/client/options"
)

const (
	// gRPC methods for querying session information
	methodQuerySession                           = "/sentinel.session.v3.QueryService/QuerySession"
	methodQuerySessions                          = "/sentinel.session.v3.QueryService/QuerySessions"
	methodQuerySessionsForAccount                = "/sentinel.session.v3.QueryService/QuerySessionsForAccount"
	methodQuerySessionsForNode                   = "/sentinel.session.v3.QueryService/QuerySessionsForNode"
	methodQuerySessionsForSubscription           = "/sentinel.session.v3.QueryService/QuerySessionsForSubscription"
	methodQuerySessionsForSubscriptionAllocation = "/sentinel.session.v3.QueryService/QuerySessionsForAllocation"
)

// Session queries and returns information about a specific session based on the provided session ID.
// It uses gRPC to send a request to the "/sentinel.session.v3.QueryService/QuerySession" endpoint.
// The result is a pointer to v3.Session and an error if the query fails.
func (c *Client) Session(ctx context.Context, id uint64, opts *options.QueryOptions) (res *v3.Session, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySessionResponse
		req  = &v3.QuerySessionRequest{
			Id: id,
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySession, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack the response and return the subscription and a nil error.
	if err := c.UnpackAny(resp.Session, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// Sessions queries and returns a list of sessions based on the provided options.
// It uses gRPC to send a request to the "/sentinel.session.v3.QueryService/QuerySessions" endpoint.
// The result is a slice of v3.Session and an error if the query fails.
func (c *Client) Sessions(ctx context.Context, opts *options.QueryOptions) (res []v3.Session, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySessionsResponse
		req  = &v3.QuerySessionsRequest{
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySessions, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack each session in the response and return the list of sessions and a nil error.
	res = make([]v3.Session, len(resp.Sessions))
	for i := 0; i < len(resp.Sessions); i++ {
		if err := c.UnpackAny(resp.Sessions[i], &res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SessionsForAccount queries and returns a list of sessions associated with a specific account
// based on the provided account address and options.
// It uses gRPC to send a request to the "/sentinel.session.v3.QueryService/QuerySessionsForAccount" endpoint.
// The result is a slice of v3.Session and an error if the query fails.
func (c *Client) SessionsForAccount(ctx context.Context, accAddr cosmossdk.AccAddress, opts *options.QueryOptions) (res []v3.Session, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySessionsForAccountResponse
		req  = &v3.QuerySessionsForAccountRequest{
			Address:    accAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySessionsForAccount, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack each session in the response and return the list of sessions and a nil error.
	res = make([]v3.Session, len(resp.Sessions))
	for i := 0; i < len(resp.Sessions); i++ {
		if err := c.UnpackAny(resp.Sessions[i], &res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SessionsForNode queries and returns a list of sessions associated with a specific node
// based on the provided node address and options.
// It uses gRPC to send a request to the "/sentinel.session.v3.QueryService/QuerySessionsForNode" endpoint.
// The result is a slice of v3.Session and an error if the query fails.
func (c *Client) SessionsForNode(ctx context.Context, nodeAddr base.NodeAddress, opts *options.QueryOptions) (res []v3.Session, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySessionsForNodeResponse
		req  = &v3.QuerySessionsForNodeRequest{
			Address:    nodeAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySessionsForNode, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack each session in the response and return the list of sessions and a nil error.
	res = make([]v3.Session, len(resp.Sessions))
	for i := 0; i < len(resp.Sessions); i++ {
		if err := c.UnpackAny(resp.Sessions[i], &res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SessionsForSubscription queries and returns a list of sessions associated with a specific subscription
// based on the provided subscription ID and options.
// It uses gRPC to send a request to the "/sentinel.session.v3.QueryService/QuerySessionsForSubscription" endpoint.
// The result is a slice of v3.Session and an error if the query fails.
func (c *Client) SessionsForSubscription(ctx context.Context, id uint64, opts *options.QueryOptions) (res []v3.Session, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySessionsForSubscriptionResponse
		req  = &v3.QuerySessionsForSubscriptionRequest{
			Id:         id,
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySessionsForSubscription, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack each session in the response and return the list of sessions and a nil error.
	res = make([]v3.Session, len(resp.Sessions))
	for i := 0; i < len(resp.Sessions); i++ {
		if err := c.UnpackAny(resp.Sessions[i], &res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SessionsForSubscriptionAllocation queries and returns a list of sessions associated with a specific subscription allocation
// based on the provided subscription ID, account address, and options.
// It uses gRPC to send a request to the "/sentinel.session.v3.QueryService/QuerySessionsForAllocation" endpoint.
// The result is a slice of v3.Session and an error if the query fails.
func (c *Client) SessionsForSubscriptionAllocation(ctx context.Context, id uint64, accAddr cosmossdk.AccAddress, opts *options.QueryOptions) (res []v3.Session, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySessionsForAllocationResponse
		req  = &v3.QuerySessionsForAllocationRequest{
			Id:         id,
			Address:    accAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySessionsForSubscriptionAllocation, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack each session in the response and return the list of sessions and a nil error.
	res = make([]v3.Session, len(resp.Sessions))
	for i := 0; i < len(resp.Sessions); i++ {
		if err := c.UnpackAny(resp.Sessions[i], &res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}
