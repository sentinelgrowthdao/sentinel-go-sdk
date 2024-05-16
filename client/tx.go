package client

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/sentinel-official/sentinel-go-sdk/v1/client/options"
	"github.com/sentinel-official/sentinel-go-sdk/v1/utils"
)

// Simulate simulates the execution of a transaction before broadcasting it.
// It takes a context, transaction bytes, and query options as input parameters,
// and returns a SimulateResponse and an error, if any.
func (c *Context) Simulate(ctx context.Context, buf []byte, opts *options.QueryOptions) (*txtypes.SimulateResponse, error) {
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
func (c *Context) simulateTx(ctx context.Context, txb client.TxBuilder, opts *options.TxOptions) (uint64, error) {
	// Encode transaction into bytes
	buf, err := c.TxEncoder()(txb.GetTx())
	if err != nil {
		return 0, err
	}

	// Simulate the transaction execution
	res, err := c.Simulate(ctx, buf, opts.QueryOptions)
	if err != nil {
		return 0, err
	}

	// Calculate gas usage
	return uint64(float64(res.GasInfo.GasUsed) * opts.GasAdjustment), nil
}

// broadcastTxSync broadcasts a transaction synchronously.
// It takes a context, a transaction builder, and transaction options as input parameters,
// and returns the broadcast result and an error, if any.
func (c *Context) broadcastTxSync(ctx context.Context, txb client.TxBuilder, opts *options.TxOptions) (*coretypes.ResultBroadcastTx, error) {
	// Encode transaction into bytes
	buf, err := c.TxEncoder()(txb.GetTx())
	if err != nil {
		return nil, err
	}

	// Get client for broadcasting
	client, err := opts.Client()
	if err != nil {
		return nil, err
	}

	// Broadcast transaction synchronously
	return client.BroadcastTxSync(ctx, buf)
}

// signTx signs a transaction with given key and account information.
// It takes a transaction builder, key information, account information, and transaction options as input parameters,
// and returns an error, if any.
func (c *Context) signTx(txb client.TxBuilder, key keyring.Info, account authtypes.AccountI, opts *options.TxOptions) error {
	// Prepare single signature data
	singleSignatureData := txsigning.SingleSignatureData{
		SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}
	// Prepare signature information
	signature := txsigning.SignatureV2{
		PubKey:   key.GetPubKey(),
		Data:     &singleSignatureData,
		Sequence: account.GetSequence(),
	}

	// Set signature in transaction builder
	if err := txb.SetSignatures(signature); err != nil {
		return err
	}

	// Prepare signer data
	signerData := authsigning.SignerData{
		ChainID:       opts.ChainID,
		AccountNumber: account.GetAccountNumber(),
		Sequence:      account.GetSequence(),
	}

	// Get sign bytes
	buf, err := c.SignModeHandler().GetSignBytes(singleSignatureData.SignMode, signerData, txb.GetTx())
	if err != nil {
		return err
	}

	// Sign transaction
	buf, _, err = c.Sign(opts.FromName, buf, opts.KeyOptions)
	if err != nil {
		return err
	}

	// Update signature data with signed bytes
	singleSignatureData = txsigning.SingleSignatureData{
		SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
		Signature: buf,
	}
	signature = txsigning.SignatureV2{
		PubKey:   key.GetPubKey(),
		Data:     &singleSignatureData,
		Sequence: account.GetSequence(),
	}

	// Set updated signature in transaction builder
	if err := txb.SetSignatures(signature); err != nil {
		return err
	}

	return nil
}

// prepareTx prepares a transaction for broadcasting.
// It takes a context, key information, account information, message(s), and transaction options as input parameters,
// and returns a transaction builder and an error, if any.
func (c *Context) prepareTx(ctx context.Context, key keyring.Info, account authtypes.AccountI, msgs []sdk.Msg, opts *options.TxOptions) (client.TxBuilder, error) {
	// Create new transaction builder
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

	// Prepare signature
	signature := txsigning.SignatureV2{
		PubKey: key.GetPubKey(),
		Data: &txsigning.SingleSignatureData{
			SignMode: txsigning.SignMode_SIGN_MODE_DIRECT,
		},
		Sequence: account.GetSequence(),
	}
	// Set signature in transaction builder
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
func (c *Context) BroadcastTx(ctx context.Context, msgs []sdk.Msg, opts *options.TxOptions) (*coretypes.ResultBroadcastTx, error) {
	// Get key for signing
	key, err := c.Key(opts.FromName, opts.KeyOptions)
	if err != nil {
		return nil, err
	}

	// Get account information
	account, err := c.Account(ctx, key.GetAddress(), opts.QueryOptions)
	if err != nil {
		return nil, err
	}

	// Prepare transaction for broadcasting
	txb, err := c.prepareTx(ctx, key, account, msgs, opts)
	if err != nil {
		return nil, err
	}

	// Sign transaction
	if err := c.signTx(txb, key, account, opts); err != nil {
		return nil, err
	}

	// Broadcast the signed transaction synchronously and return the result.
	return c.broadcastTxSync(ctx, txb, opts)
}
