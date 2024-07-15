package client

import (
	"context"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"

	"github.com/sentinel-official/sentinel-go-sdk/v1/client/options"
)

const (
	// gRPC methods for querying subscription information
	methodQuerySubscription            = "/sentinel.subscription.v2.QueryService/QuerySubscription"
	methodQuerySubscriptions           = "/sentinel.subscription.v2.QueryService/QuerySubscriptions"
	methodQuerySubscriptionsForAccount = "/sentinel.subscription.v2.QueryService/QuerySubscriptionsForAccount"
	methodQuerySubscriptionsForNode    = "/sentinel.subscription.v2.QueryService/QuerySubscriptionsForNode"
	methodQuerySubscriptionsForPlan    = "/sentinel.subscription.v2.QueryService/QuerySubscriptionsForPlan"

	// gRPC methods for querying subscription allocation information
	methodQuerySubscriptionAllocation  = "/sentinel.subscription.v2.QueryService/QueryAllocation"
	methodQuerySubscriptionAllocations = "/sentinel.subscription.v2.QueryService/QueryAllocations"

	// gRPC methods for querying subscription payout information
	methodQuerySubscriptionPayout            = "/sentinel.subscription.v2.QueryService/QueryPayout"
	methodQuerySubscriptionPayouts           = "/sentinel.subscription.v2.QueryService/QueryPayouts"
	methodQuerySubscriptionPayoutsForAccount = "/sentinel.subscription.v2.QueryService/QueryPayoutsForAccount"
	methodQuerySubscriptionPayoutsForNode    = "/sentinel.subscription.v2.QueryService/QueryPayoutsForNode"
)

// Subscription queries and returns information about a specific subscription based on the provided subscription ID.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QuerySubscription" endpoint.
// The result is a v2.Subscription and an error if the query fails.
func (c *Context) Subscription(ctx context.Context, id uint64, opts *options.QueryOptions) (res v2.Subscription, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QuerySubscriptionResponse
		req  = &v2.QuerySubscriptionRequest{
			Id: id,
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscription, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack the response and return the subscription and a nil error.
	if err := c.UnpackAny(resp.Subscription, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// Subscriptions queries and returns a list of subscriptions based on the provided options.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QuerySubscriptions" endpoint.
// The result is a slice of v2.Subscription and an error if the query fails.
func (c *Context) Subscriptions(ctx context.Context, opts *options.QueryOptions) (res []v2.Subscription, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QuerySubscriptionsResponse
		req  = &v2.QuerySubscriptionsRequest{
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptions, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack each subscription in the response and return the list of subscriptions and a nil error.
	res = make([]v2.Subscription, len(resp.Subscriptions))
	for i := 0; i < len(resp.Subscriptions); i++ {
		if err := c.UnpackAny(resp.Subscriptions[i], &res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SubscriptionsForAccount queries and returns a list of subscriptions associated with a specific account.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QuerySubscriptionsForAccount" endpoint.
// The result is a slice of v2.Subscription and an error if the query fails.
// The account is identified by the provided cosmossdk.AccAddress.
func (c *Context) SubscriptionsForAccount(ctx context.Context, accAddr cosmossdk.AccAddress, opts *options.QueryOptions) (res []v2.Subscription, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QuerySubscriptionsForAccountResponse
		req  = &v2.QuerySubscriptionsForAccountRequest{
			Address:    accAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionsForAccount, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack each subscription in the response and return the list of subscriptions and a nil error.
	res = make([]v2.Subscription, len(resp.Subscriptions))
	for i := 0; i < len(resp.Subscriptions); i++ {
		if err := c.UnpackAny(resp.Subscriptions[i], &res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SubscriptionsForNode queries and returns a list of subscriptions associated with a specific node.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QuerySubscriptionsForNode" endpoint.
// The result is a slice of v2.Subscription and an error if the query fails.
// The node is identified by the provided base.NodeAddress.
func (c *Context) SubscriptionsForNode(ctx context.Context, nodeAddr base.NodeAddress, opts *options.QueryOptions) (res []v2.Subscription, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QuerySubscriptionsForNodeResponse
		req  = &v2.QuerySubscriptionsForNodeRequest{
			Address:    nodeAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionsForNode, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack each subscription in the response and return the list of subscriptions and a nil error.
	res = make([]v2.Subscription, len(resp.Subscriptions))
	for i := 0; i < len(resp.Subscriptions); i++ {
		if err := c.UnpackAny(resp.Subscriptions[i], &res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SubscriptionsForPlan queries and returns a list of subscriptions associated with a specific plan.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QuerySubscriptionsForPlan" endpoint.
// The result is a slice of v2.Subscription and an error if the query fails.
// The plan is identified by the provided ID.
func (c *Context) SubscriptionsForPlan(ctx context.Context, id uint64, opts *options.QueryOptions) (res []v2.Subscription, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QuerySubscriptionsForPlanResponse
		req  = &v2.QuerySubscriptionsForPlanRequest{
			Id:         id,
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionsForPlan, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack each subscription in the response and return the list of subscriptions and a nil error.
	res = make([]v2.Subscription, len(resp.Subscriptions))
	for i := 0; i < len(resp.Subscriptions); i++ {
		if err := c.UnpackAny(resp.Subscriptions[i], &res[i]); err != nil {
			return nil, err
		}
	}

	return res, nil
}

// SubscriptionAllocation queries and returns information about a specific allocation within a subscription.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QueryAllocation" endpoint.
// The result is a pointer to v2.Allocation and an error if the query fails.
func (c *Context) SubscriptionAllocation(ctx context.Context, id uint64, accAddr cosmossdk.AccAddress, opts *options.QueryOptions) (res *v2.Allocation, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryAllocationResponse
		req  = &v2.QueryAllocationRequest{
			Id:      id,
			Address: accAddr.String(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionAllocation, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return a pointer to the allocation and a nil error.
	return &resp.Allocation, nil
}

// SubscriptionAllocations queries and returns a list of allocations within a specific subscription.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QueryAllocations" endpoint.
// The result is a slice of v2.Allocation and an error if the query fails.
func (c *Context) SubscriptionAllocations(ctx context.Context, id uint64, opts *options.QueryOptions) (res []v2.Allocation, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryAllocationsResponse
		req  = &v2.QueryAllocationsRequest{
			Id:         id,
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionAllocations, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of allocations and a nil error.
	return resp.Allocations, nil
}

// SubscriptionPayout queries and returns information about a specific payout within a subscription.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QueryPayout" endpoint.
// The result is a pointer to v2.Payout and an error if the query fails.
func (c *Context) SubscriptionPayout(ctx context.Context, id uint64, opts *options.QueryOptions) (res *v2.Payout, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryPayoutResponse
		req  = &v2.QueryPayoutRequest{
			Id: id,
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionPayout, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return a pointer to the payout and a nil error.
	return &resp.Payout, nil
}

// SubscriptionPayouts queries and returns a list of payouts within a specific subscription.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QueryPayouts" endpoint.
// The result is a slice of v2.Payout and an error if the query fails.
func (c *Context) SubscriptionPayouts(ctx context.Context, opts *options.QueryOptions) (res []v2.Payout, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryPayoutsResponse
		req  = &v2.QueryPayoutsRequest{
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionPayouts, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of payouts and a nil error.
	return resp.Payouts, nil
}

// SubscriptionPayoutsForAccount queries and returns a list of payouts associated with a specific account.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QueryPayoutsForAccount" endpoint.
// The result is a slice of v2.Payout and an error if the query fails.
// The account is identified by the provided cosmossdk.AccAddress.
func (c *Context) SubscriptionPayoutsForAccount(ctx context.Context, accAddr cosmossdk.AccAddress, opts *options.QueryOptions) (res []v2.Payout, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryPayoutsForAccountResponse
		req  = &v2.QueryPayoutsForAccountRequest{
			Address:    accAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionPayoutsForAccount, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of payouts and a nil error.
	return resp.Payouts, nil
}

// SubscriptionPayoutsForNode queries and returns a list of payouts associated with a specific node.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QueryPayoutsForNode" endpoint.
// The result is a slice of v2.Payout and an error if the query fails.
// The node is identified by the provided base.NodeAddress.
func (c *Context) SubscriptionPayoutsForNode(ctx context.Context, nodeAddr base.NodeAddress, opts *options.QueryOptions) (res []v2.Payout, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryPayoutsForNodeResponse
		req  = &v2.QueryPayoutsForNodeRequest{
			Address:    nodeAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionPayoutsForNode, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of payouts and a nil error.
	return resp.Payouts, nil
}
