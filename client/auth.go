package client

import (
	"context"

	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	// gRPC methods for querying account information
	methodQueryAccount  = "/cosmos.auth.v1beta1.Query/Account"
	methodQueryAccounts = "/cosmos.auth.v1beta1.Query/Accounts"
)

// Account queries and returns an account using the given address and options.
// It uses gRPC to send a request to the "/cosmos.auth.v1beta1.Query/Account" endpoint.
// The result is an authtypes.AccountI interface and an error if the query fails.
func (c *Client) Account(ctx context.Context, accAddr cosmossdk.AccAddress, opts *Options) (res authtypes.AccountI, err error) {
	// Initialize variables for the query.
	var (
		resp authtypes.QueryAccountResponse
		req  = &authtypes.QueryAccountRequest{
			Address: accAddr.String(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryAccount, req, &resp, opts); err != nil {
		return nil, err
	}

	// Unpack the Any type account from the response.
	if err := c.UnpackAny(resp.Account, &res); err != nil {
		return nil, err
	}

	// Return the account and a nil error.
	return res, nil
}

// Accounts queries and returns a list of accounts using the given options.
// It uses gRPC to send a request to the "/cosmos.auth.v1beta1.Query/Accounts" endpoint.
// The result is a slice of authtypes.AccountI and an error if the query fails.
func (c *Client) Accounts(ctx context.Context, opts *Options) (res []authtypes.AccountI, err error) {
	// Initialize variables for the query.
	var (
		resp authtypes.QueryAccountsResponse
		req  = &authtypes.QueryAccountsRequest{
			Pagination: opts.PageRequest(),
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, methodQueryAccounts, req, &resp, opts); err != nil {
		return nil, err
	}

	// Initialize a slice to store the accounts.
	res = make([]authtypes.AccountI, len(resp.Accounts))

	// Unpack each Any type account from the response and add it to the result slice.
	for i := 0; i < len(resp.Accounts); i++ {
		if err := c.UnpackAny(resp.Accounts[i], &res[i]); err != nil {
			return nil, err
		}
	}

	// Return the list of accounts and a nil error.
	return res, nil
}
