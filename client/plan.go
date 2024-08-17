package client

import (
	"context"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
)

const (
	// gRPC methods for querying plan information
	methodQueryPlan             = "/sentinel.plan.v2.QueryService/QueryPlan"
	methodQueryPlans            = "/sentinel.plan.v2.QueryService/QueryPlans"
	methodQueryPlansForProvider = "/sentinel.plan.v2.QueryService/QueryPlansForProvider"
)

// Plan queries and returns information about a specific plan based on the provided plan ID.
// It uses gRPC to send a request to the "/sentinel.plan.v2.QueryService/QueryPlan" endpoint.
// The result is a pointer to v2.Plan and an error if the query fails.
func (c *Client) Plan(ctx context.Context, id uint64, opts *Options) (res *v2.Plan, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryPlanResponse
		req  = &v2.QueryPlanRequest{
			Id: id,
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryPlan, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return a pointer to the plan and a nil error.
	return &resp.Plan, nil
}

// Plans queries and returns a list of plans based on the provided status and options.
// It uses gRPC to send a request to the "/sentinel.plan.v2.QueryService/QueryPlans" endpoint.
// The result is a slice of v2.Plan and an error if the query fails.
func (c *Client) Plans(ctx context.Context, status v1base.Status, opts *Options) (res []v2.Plan, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryPlansResponse
		req  = &v2.QueryPlansRequest{
			Status:     status,
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryPlans, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of plans and a nil error.
	return resp.Plans, nil
}

// PlansForProvider queries and returns a list of plans associated with a specific provider
// based on the provided provider address, status, and options.
// It uses gRPC to send a request to the "/sentinel.plan.v2.QueryService/QueryPlansForProvider" endpoint.
// The result is a slice of v2.Plan and an error if the query fails.
func (c *Client) PlansForProvider(ctx context.Context, provAddr base.ProvAddress, status v1base.Status, opts *Options) (res []v2.Plan, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryPlansForProviderResponse
		req  = &v2.QueryPlansForProviderRequest{
			Address:    provAddr.String(),
			Status:     status,
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryPlansForProvider, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of plans and a nil error.
	return resp.Plans, nil
}
