package client

import (
	"context"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"

	"github.com/sentinel-official/sentinel-go-sdk/client/options"
)

const (
	// gRPC methods for querying subscription information
	methodQuerySubscription            = "/sentinel.subscription.v2.QueryService/QuerySubscription"
	methodQuerySubscriptions           = "/sentinel.subscription.v2.QueryService/QuerySubscriptions"
	methodQuerySubscriptionsForAccount = "/sentinel.subscription.v2.QueryService/QuerySubscriptionsForAccount"
	methodQuerySubscriptionsForPlan    = "/sentinel.subscription.v2.QueryService/QuerySubscriptionsForPlan"

	// gRPC methods for querying subscription allocation information
	methodQuerySubscriptionAllocation  = "/sentinel.subscription.v2.QueryService/QueryAllocation"
	methodQuerySubscriptionAllocations = "/sentinel.subscription.v2.QueryService/QueryAllocations"
)

// Subscription queries and returns information about a specific subscription based on the provided subscription ID.
// It uses gRPC to send a request to the "/sentinel.subscription.v3.QueryService/QuerySubscription" endpoint.
// The result is a v3.Subscription and an error if the query fails.
func (c *Client) Subscription(ctx context.Context, id uint64, opts *options.Options) (res *v3.Subscription, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySubscriptionResponse
		req  = &v3.QuerySubscriptionRequest{
			Id: id,
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscription, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return a pointer to the subscription and a nil error.
	return &resp.Subscription, nil
}

// Subscriptions queries and returns a list of subscriptions based on the provided options.
// It uses gRPC to send a request to the "/sentinel.subscription.v3.QueryService/QuerySubscriptions" endpoint.
// The result is a slice of v3.Subscription and an error if the query fails.
func (c *Client) Subscriptions(ctx context.Context, opts *options.Options) (res []v3.Subscription, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySubscriptionsResponse
		req  = &v3.QuerySubscriptionsRequest{
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptions, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of subscriptions and a nil error.
	return resp.Subscriptions, nil
}

// SubscriptionsForAccount queries and returns a list of subscriptions associated with a specific account.
// It uses gRPC to send a request to the "/sentinel.subscription.v3.QueryService/QuerySubscriptionsForAccount" endpoint.
// The result is a slice of v3.Subscription and an error if the query fails.
// The account is identified by the provided cosmossdk.AccAddress.
func (c *Client) SubscriptionsForAccount(ctx context.Context, accAddr cosmossdk.AccAddress, opts *options.Options) (res []v3.Subscription, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySubscriptionsForAccountResponse
		req  = &v3.QuerySubscriptionsForAccountRequest{
			Address:    accAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionsForAccount, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of subscriptions and a nil error.
	return resp.Subscriptions, nil
}

// SubscriptionsForPlan queries and returns a list of subscriptions associated with a specific plan.
// It uses gRPC to send a request to the "/sentinel.subscription.v3.QueryService/QuerySubscriptionsForPlan" endpoint.
// The result is a slice of v3.Subscription and an error if the query fails.
// The plan is identified by the provided ID.
func (c *Client) SubscriptionsForPlan(ctx context.Context, id uint64, opts *options.Options) (res []v3.Subscription, err error) {
	// Initialize variables for the query.
	var (
		resp v3.QuerySubscriptionsForPlanResponse
		req  = &v3.QuerySubscriptionsForPlanRequest{
			Id:         id,
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQuerySubscriptionsForPlan, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of subscriptions and a nil error.
	return resp.Subscriptions, nil
}

// SubscriptionAllocation queries and returns information about a specific allocation within a subscription.
// It uses gRPC to send a request to the "/sentinel.subscription.v2.QueryService/QueryAllocation" endpoint.
// The result is a pointer to v2.Allocation and an error if the query fails.
func (c *Client) SubscriptionAllocation(ctx context.Context, id uint64, accAddr cosmossdk.AccAddress, opts *options.Options) (res *v2.Allocation, err error) {
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
func (c *Client) SubscriptionAllocations(ctx context.Context, id uint64, opts *options.Options) (res []v2.Allocation, err error) {
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
