package client

import (
	"context"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"

	"github.com/sentinel-official/sentinel-go-sdk/v1/client/options"
)

const (
	// gRPC methods for querying provider information
	methodQueryProvider  = "/sentinel.provider.v2.QueryService/QueryProvider"
	methodQueryProviders = "/sentinel.provider.v2.QueryService/QueryProviders"
)

// Provider queries and returns information about a specific provider based on the provided provider address.
// It uses gRPC to send a request to the "/sentinel.provider.v2.QueryService/QueryProvider" endpoint.
// The result is a pointer to v2.Provider and an error if the query fails.
func (c *Client) Provider(ctx context.Context, provAddr base.ProvAddress, opts *options.QueryOptions) (res *v2.Provider, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryProviderResponse
		req  = &v2.QueryProviderRequest{
			Address: provAddr.String(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryProvider, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return a pointer to the provider and a nil error.
	return &resp.Provider, nil
}

// Providers queries and returns a list of providers based on the provided status and options.
// It uses gRPC to send a request to the "/sentinel.provider.v2.QueryService/QueryProviders" endpoint.
// The result is a slice of v2.Provider and an error if the query fails.
func (c *Client) Providers(ctx context.Context, status v1base.Status, opts *options.QueryOptions) (res []v2.Provider, err error) {
	// Initialize variables for the query.
	var (
		resp v2.QueryProvidersResponse
		req  = &v2.QueryProvidersRequest{
			Status:     status,
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryProviders, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of providers and a nil error.
	return resp.Providers, nil
}
