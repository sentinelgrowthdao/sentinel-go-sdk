package client

import (
	"context"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"

	"github.com/sentinel-official/sentinel-go-sdk/v1/client/options"
)

const (
	// gRPC methods for querying node information
	methodQueryNode         = "/sentinel.node.v2.QueryService/QueryNode"
	methodQueryNodes        = "/sentinel.node.v2.QueryService/QueryNodes"
	methodQueryNodesForPlan = "/sentinel.node.v2.QueryService/QueryNodesForPlan"
)

// Node queries and returns information about a specific node based on the provided node address.
// It uses gRPC to send a request to the "/sentinel.node.v2.QueryService/QueryNode" endpoint.
// The result is a pointer to nodetypes.Node and an error if the query fails.
func (c *Context) Node(ctx context.Context, nodeAddr base.NodeAddress, opts *options.QueryOptions) (res *nodetypes.Node, err error) {
	// Initialize variables for the query.
	var (
		resp nodetypes.QueryNodeResponse
		req  = &nodetypes.QueryNodeRequest{
			Address: nodeAddr.String(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryNode, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return a pointer to the node and a nil error.
	return &resp.Node, nil
}

// Nodes queries and returns a list of nodes based on the provided status and options.
// It uses gRPC to send a request to the "/sentinel.node.v2.QueryService/QueryNodes" endpoint.
// The result is a slice of nodetypes.Node and an error if the query fails.
func (c *Context) Nodes(ctx context.Context, status v1base.Status, opts *options.QueryOptions) (res []nodetypes.Node, err error) {
	// Initialize variables for the query.
	var (
		resp nodetypes.QueryNodesResponse
		req  = &nodetypes.QueryNodesRequest{
			Status:     status,
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryNodes, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of nodes and a nil error.
	return resp.Nodes, nil
}

// NodesForPlan queries and returns a list of nodes associated with a specific plan
// based on the provided plan ID, status, and options.
// It uses gRPC to send a request to the "/sentinel.node.v2.QueryService/QueryNodesForPlan" endpoint.
// The result is a slice of nodetypes.Node and an error if the query fails.
func (c *Context) NodesForPlan(ctx context.Context, id uint64, status v1base.Status, opts *options.QueryOptions) (res []nodetypes.Node, err error) {
	// Initialize variables for the query.
	var (
		resp nodetypes.QueryNodesForPlanResponse
		req  = &nodetypes.QueryNodesForPlanRequest{
			Id:         id,
			Status:     status,
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryNodesForPlan, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the list of nodes and a nil error.
	return resp.Nodes, nil
}
