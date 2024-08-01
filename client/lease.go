package client

import (
	"context"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"

	"github.com/sentinel-official/sentinel-go-sdk/client/options"
)

const (
	// gRPC methods for querying lease information
	methodQueryLease             = "/sentinel.lease.v1.QueryService/QueryLease"
	methodQueryLeases            = "/sentinel.lease.v1.QueryService/QueryLeases"
	methodQueryLeasesForNode     = "/sentinel.lease.v1.QueryService/QueryLeasesForNode"
	methodQueryLeasesForProvider = "/sentinel.lease.v1.QueryService/QueryLeasesForProvider"
)

// Lease queries and returns information about a specific lease based on the provided lease ID.
// It uses gRPC to send a request to the "/sentinel.lease.v1.QueryService/QueryLease" endpoint.
// The result is a pointer to v1.Lease and an error if the query fails.
func (c *Client) Lease(ctx context.Context, id uint64, opts *options.Options) (res *v1.Lease, err error) {
	// Initialize variables for the query.
	var (
		resp v1.QueryLeaseResponse
		req  = &v1.QueryLeaseRequest{
			Id: id,
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryLease, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return a pointer to the lease and a nil error.
	return &resp.Lease, nil
}

// Leases queries and returns a list of leases based on the provided options.
// It uses gRPC to send a request to the "/sentinel.lease.v1.QueryService/QueryLeases" endpoint.
// The result is a slice of v1.Lease and an error if the query fails.
func (c *Client) Leases(ctx context.Context, opts *options.Options) (res []v1.Lease, err error) {
	// Initialize variables for the query.
	var (
		resp v1.QueryLeasesResponse
		req  = &v1.QueryLeasesRequest{
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryLeases, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of leases and a nil error.
	return resp.Leases, nil
}

// LeasesForNode queries and returns a list of leases associated with a specific node.
// It uses gRPC to send a request to the "/sentinel.lease.v1.QueryService/QueryLeasesForNode" endpoint.
// The result is a slice of v1.Lease and an error if the query fails.
// The node is identified by the provided base.NodeAddress.
func (c *Client) LeasesForNode(ctx context.Context, nodeAddr base.NodeAddress, opts *options.Options) (res []v1.Lease, err error) {
	// Initialize variables for the query.
	var (
		resp v1.QueryLeasesForNodeResponse
		req  = &v1.QueryLeasesForNodeRequest{
			Address:    nodeAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryLeasesForNode, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of leases and a nil error.
	return resp.Leases, nil
}

// LeasesForProvider queries and returns a list of leases associated with a specific provider.
// It uses gRPC to send a request to the "/sentinel.lease.v1.QueryService/QueryLeasesForProvider" endpoint.
// The result is a slice of v1.Lease and an error if the query fails.
// The provider is identified by the provided base.ProvAddress.
func (c *Client) LeasesForProvider(ctx context.Context, provAddr base.ProvAddress, opts *options.Options) (res []v1.Lease, err error) {
	// Initialize variables for the query.
	var (
		resp v1.QueryLeasesForProviderResponse
		req  = &v1.QueryLeasesForProviderRequest{
			Address:    provAddr.String(),
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryLeasesForProvider, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of leases and a nil error.
	return resp.Leases, nil
}
