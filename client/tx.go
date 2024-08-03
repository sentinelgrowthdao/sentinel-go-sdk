package client

import (
	"context"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/sentinel-official/sentinel-go-sdk/client/options"
	"github.com/sentinel-official/sentinel-go-sdk/utils"
)

// Simulate simulates the execution of a transaction before broadcasting it.
// It takes a context, transaction bytes, and query options as input parameters,
// and returns a SimulateResponse and an error, if any.
func (c *Client) Simulate(ctx context.Context, buf []byte, opts *options.Options) (*txtypes.SimulateResponse, error) {
	// Initialize variables for the query.
	var (
		resp   txtypes.SimulateResponse
		method = "/cosmos.tx.v1beta1.Service/Simulate"
		req    = &txtypes.SimulateRequest{
			TxBytes: buf,
		}
	)

	// Send a gRPC query using the provided context, method, request, response, and options.
	if err := c.QueryGRPC(ctx, method, req, &resp, opts); err != nil {
		return nil, err
	}

	// Return the simulation response and a nil error.
	return &resp, nil
}

// simulateTx simulates the gas usage of a transaction.
// It takes a context, a transaction builder, and transaction options as input parameters,
// and returns the gas usage and an error, if any.
func (c *Client) simulateTx(ctx context.Context, txb client.TxBuilder, opts *options.Options) (uint64, error) {
	// Encode transaction into bytes
	buf, err := c.TxEncoder()(txb.GetTx())
	if err != nil {
		return 0, err
	}

	// Simulate the transaction execution
	res, err := c.Simulate(ctx, buf, opts)
	if err != nil {
		return 0, err
	}

	// Calculate gas usage
	return uint64(opts.GasAdjustment * float64(res.GasInfo.GasUsed)), nil
}

// broadcastTxSync broadcasts a transaction synchronously.
// It takes a context, a transaction builder, and transaction options as input parameters,
// and returns the broadcast result and an error, if any.
func (c *Client) broadcastTxSync(ctx context.Context, txb client.TxBuilder, opts *options.Options) (*coretypes.ResultBroadcastTx, error) {
	// Encode transaction into bytes
	buf, err := c.TxEncoder()(txb.GetTx())
	if err != nil {
		return nil, err
	}

	// Get client for broadcasting
	rpc, err := opts.Client()
	if err != nil {
		return nil, err
	}

	// Broadcast transaction synchronously
	return rpc.BroadcastTxSync(ctx, buf)
}

// signTx signs a transaction with given key and account information.
// It takes a transaction builder, key information, account information, and transaction options as input parameters,
// and returns an error, if any.
func (c *Client) signTx(txb client.TxBuilder, key *keyring.Record, account authtypes.AccountI, opts *options.Options) error {
	// Prepare single signature data
	singleSignatureData := txsigning.SingleSignatureData{
		SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}

	// Retrieve the public key from the key record
	pubKey, err := key.GetPubKey()
	if err != nil {
		return err
	}

	// Prepare signature information
	signature := txsigning.SignatureV2{
		PubKey:   pubKey,
		Data:     &singleSignatureData,
		Sequence: account.GetSequence(),
	}

	// Set the initial empty signature in the transaction builder
	if err := txb.SetSignatures(signature); err != nil {
		return err
	}

	// Prepare signer data for creating the sign bytes
	signerData := authsigning.SignerData{
		ChainID:       opts.ChainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}

	// Get the bytes to sign from the transaction builder
	buf, err := c.SignModeHandler().GetSignBytes(singleSignatureData.SignMode, signerData, txb.GetTx())
	if err != nil {
		return err
	}

	// Sign the transaction bytes
	buf, _, err = c.Sign(opts.FromName, buf, opts)
	if err != nil {
		return err
	}

	// Update the signature data with the actual signature bytes
	singleSignatureData = txsigning.SingleSignatureData{
		SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
		Signature: buf,
	}
	signature = txsigning.SignatureV2{
		PubKey:   pubKey,
		Data:     &singleSignatureData,
		Sequence: account.GetSequence(),
	}

	// Set the updated signature in the transaction builder
	if err := txb.SetSignatures(signature); err != nil {
		return err
	}

	return nil
}

// prepareTx prepares a transaction for broadcasting.
// It takes a context, key information, account information, message(s), and transaction options as input parameters,
// and returns a transaction builder and an error, if any.
func (c *Client) prepareTx(ctx context.Context, key *keyring.Record, account authtypes.AccountI, msgs []sdk.Msg, opts *options.Options) (client.TxBuilder, error) {
	// Create a new transaction builder instance
	txb := c.NewTxBuilder()
	if err := txb.SetMsgs(msgs...); err != nil {
		return nil, err
	}

	// Set transaction fee, fee granter, gas limit, memo, and timeout height
	txb.SetFeeAmount(nil)
	txb.SetFeeGranter(utils.MustAccAddrFromBech32(opts.FeeGranterAddr))
	txb.SetGasLimit(opts.Gas)
	txb.SetMemo(opts.Memo)
	txb.SetTimeoutHeight(opts.TimeoutHeight)

	// Retrieve the public key from the key record
	pubKey, err := key.GetPubKey()
	if err != nil {
		return nil, err
	}

	// Prepare the signature data
	signature := txsigning.SignatureV2{
		PubKey: pubKey,
		Data: &txsigning.SingleSignatureData{
			SignMode: txsigning.SignMode_SIGN_MODE_DIRECT,
		},
		Sequence: account.GetSequence(),
	}

	// Set the initial empty signature in the transaction builder
	if err := txb.SetSignatures(signature); err != nil {
		return nil, err
	}

	// If set to simulate and execute, calculate gas usage and update gas limit
	if opts.SimulateAndExecute {
		gasLimit, err := c.simulateTx(ctx, txb, opts)
		if err != nil {
			return nil, err
		}

		txb.SetGasLimit(gasLimit)
	}

	return txb, nil
}

// BroadcastTx broadcasts a signed transaction.
// It takes a context, message(s), and transaction options as input parameters,
// and returns the broadcast result and an error, if any.
func (c *Client) BroadcastTx(ctx context.Context, msgs []sdk.Msg, opts *options.Options) (*coretypes.ResultBroadcastTx, error) {
	// Get key for signing
	key, err := c.Key(opts.FromName, opts)
	if err != nil {
		return nil, err
	}

	// Retrieve the address from the key record
	accAddr, err := key.GetAddress()
	if err != nil {
		return nil, err
	}

	// Get account information for the address
	account, err := c.Account(ctx, accAddr, opts)
	if err != nil {
		return nil, err
	}

	// Prepare the transaction for broadcasting
	txb, err := c.prepareTx(ctx, key, account, msgs, opts)
	if err != nil {
		return nil, err
	}

	// Sign the transaction
	if err := c.signTx(txb, key, account, opts); err != nil {
		return nil, err
	}

	// Broadcast the signed transaction synchronously and return the result.
	return c.broadcastTxSync(ctx, txb, opts)
}

// Tx retrieves a transaction from the blockchain using its hash.
// It takes a context, a transaction hash, and query options as input parameters,
// and returns the transaction result and an error, if any.
func (c *Client) Tx(ctx context.Context, hash []byte, opts *options.Options) (*coretypes.ResultTx, error) {
	// Get client for querying the blockchain
	rpc, err := opts.Client()
	if err != nil {
		return nil, err
	}

	// Perform the blockchain query for the transaction
	return rpc.Tx(ctx, hash, opts.Prove)
}
